package debugger

import (
	"testing"
)

func TestDebuggerInitialization(t *testing.T) {
    dbg := NewDebugger()
    
    if dbg.cpu == nil {
        t.Error("CPU not initialized")
    }
    
    if dbg.memory == nil {
        t.Error("Memory not initialized")
    }
    
    if dbg.breakpoints == nil {
        t.Error("Breakpoints not initialized")
    }
    
    if dbg.symbols == nil {
        t.Error("Symbols not initialized")
    }
}

func TestDebuggerBreakpoints(t *testing.T) {
    dbg := NewDebugger()
    
    // Test setting breakpoints
    dbg.SetBreakpoint(0x1000)
    if !dbg.IsBreakpointEnabled(0x1000) {
        t.Error("Breakpoint not set correctly")
    }
    
    // Test disabling breakpoints
    dbg.DisableBreakpoint(0x1000)
    if dbg.IsBreakpointEnabled(0x1000) {
        t.Error("Breakpoint not disabled")
    }
    
    // Test enabling breakpoints
    dbg.EnableBreakpoint(0x1000)
    if !dbg.IsBreakpointEnabled(0x1000) {
        t.Error("Breakpoint not enabled")
    }
    
    // Test removing breakpoints
    dbg.RemoveBreakpoint(0x1000)
    if dbg.IsBreakpointEnabled(0x1000) {
        t.Error("Breakpoint not removed")
    }
}

func TestDebuggerTracing(t *testing.T) {
    dbg := NewDebugger()
    
    dbg.SetTracing(true)
    if !dbg.tracing {
        t.Error("Tracing not enabled")
    }
    
    dbg.SetTracing(false)
    if dbg.tracing {
        t.Error("Tracing not disabled")
    }
}