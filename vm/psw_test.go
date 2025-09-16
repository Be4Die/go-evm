package vm

import "testing"

func TestPSWFlags(t *testing.T) {
    psw := NewPSW()
    
    // Test setting and getting flags
    psw.SetFlag(FLAG_ZERO, true)
    if !psw.GetFlag(FLAG_ZERO) {
        t.Error("Zero flag not set correctly")
    }

    psw.SetFlag(FLAG_CARRY, false)
    if psw.GetFlag(FLAG_CARRY) {
        t.Error("Carry flag not cleared correctly")
    }
}

func TestPSWRegisters(t *testing.T) {
    psw := NewPSW()
    
    psw.SetIP(0xABCD)
    if psw.GetIP() != 0xABCD {
        t.Error("IP not set correctly")
    }

    psw.SetSP(0x1F)
    if psw.GetSP() != 0x1F {
        t.Error("SP not set correctly")
    }
}