package state

import (
	"sync"
	"time"

	"github.com/kryptamine/herdr-auto-title/internal/herdr"
)

// Changes remembers what a snapshot cannot say: when each pane last changed.
// Pane revisions are monotonic, so comparing one poll's with the last says
// which panes drew.
type Changes struct {
	mu    sync.Mutex
	panes map[string]paneChange
	// focus is the pane last seen focused in each tab. Herdr flags only the
	// viewed tab's pane, so a tab you leave keeps the one you were in.
	focus map[string]string
	now   func() time.Time
}

type paneChange struct {
	revision uint64
	at       time.Time
}

func NewChanges() *Changes {
	return &Changes{
		panes: make(map[string]paneChange),
		focus: make(map[string]string),
		now:   time.Now,
	}
}

// Observe records a poll: panes whose revision moved changed just now, and
// panes the session no longer holds are forgotten. It returns the panes that
// changed, a first sighting included.
func (c *Changes) Observe(panes []herdr.PaneInfo) map[string]bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := c.now()

	seen := make(map[string]paneChange, len(panes))
	drew := make(map[string]bool)

	for _, pane := range panes {
		previous, known := c.panes[pane.PaneID]
		// Any difference, not just an advance: a revision that went backwards
		// is a new pane wearing an id Herdr has handed out again.
		switch {
		case !known, pane.Revision != previous.revision:
			seen[pane.PaneID] = paneChange{revision: pane.Revision, at: now}
			drew[pane.PaneID] = true
		default:
			seen[pane.PaneID] = previous
		}
	}

	c.panes = seen

	focus := make(map[string]string, len(c.focus))

	for tab, paneID := range c.focus {
		if _, ok := seen[paneID]; ok {
			focus[tab] = paneID
		}
	}

	for _, pane := range panes {
		if pane.Focused {
			focus[pane.TabID] = pane.PaneID
		}
	}

	c.focus = focus

	return drew
}

// LastFocused reports the pane last seen focused in a tab, or "" for a tab no
// poll has seen focused yet.
func (c *Changes) LastFocused(tabID string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.focus[tabID]
}

// ChangedAt reports when a pane was last seen to change, or the zero time for
// a pane no poll has covered yet.
func (c *Changes) ChangedAt(paneID string) time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.panes[paneID].at
}
