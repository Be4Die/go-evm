package debugger

import "testing"

func TestBreakpointManager(t *testing.T) {
    bm := NewBreakpointManager()

    // Test SetBreakpoint
    bm.SetBreakpoint(0x1000)
    if !bm.IsBreakpoint(0x1000) {
        t.Error("Breakpoint not set")
    }

    // Test IsEnabled (should be enabled by default)
    if !bm.IsEnabled(0x1000) {
        t.Error("Breakpoint should be enabled by default")
    }

    // Test DisableBreakpoint
    bm.DisableBreakpoint(0x1000)
    if bm.IsEnabled(0x1000) {
        t.Error("Breakpoint should be disabled")
    }

    // Test EnableBreakpoint
    bm.EnableBreakpoint(0x1000)
    if !bm.IsEnabled(0x1000) {
        t.Error("Breakpoint should be enabled")
    }

    // Test RemoveBreakpoint
    bm.RemoveBreakpoint(0x1000)
    if bm.IsBreakpoint(0x1000) {
        t.Error("Breakpoint should be removed")
    }

    // Test GetBreakpoints
    bm.SetBreakpoint(0x2000)
    bm.SetBreakpoint(0x3000)
    breakpoints := bm.GetBreakpoints()
    if len(breakpoints) != 2 {
        t.Errorf("Expected 2 breakpoints, got %d", len(breakpoints))
    }
}