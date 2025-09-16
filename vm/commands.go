// vm/commands.go
package vm

import (
	"errors"
	"fmt"
	"math"
)

func (c *CPU) push(value uint32) error {
	if c.psw.sp == 0 {
		return errors.New("stack overflow")
	}
	c.psw.sp--
	stackAddr := c.getStackAddress()
	return c.memory.WriteWordAt(stackAddr, value)
}

func (c *CPU) pop() (uint32, error) {
	stackAddr := c.getStackAddress()
	value, err := c.memory.ReadWordAt(stackAddr)
	if err != nil {
		return 0, err
	}
	c.psw.sp++
	if c.psw.sp > 31 {
		c.psw.sp = 31 // Reset to top if we exceed
	}
	return value, nil
}

func (c *CPU) getStackAddress() uint16 {
	return uint16(31-c.psw.sp) * 4 // 32 words of 4 bytes each
}

func (c *CPU) readImmediateOffset() (uint16, error) {
    low, err := c.memory.ReadByteAt(c.psw.ip)
    if err != nil {
        return 0, err
    }
    high, err := c.memory.ReadByteAt(c.psw.ip + 1)
    if err != nil {
        return 0, err
    }
    c.psw.ip += 2
    return uint16(low) | (uint16(high) << 8), nil
}

// Command implementations
func (c *CPU) movCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Pop value from stack
	value, err := c.pop()
	if err != nil {
		return err
	}

	// Write to memory at offset
	return c.memory.WriteWordAt(offset, value)
}

func (c *CPU) addIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform addition
	result := stackValue + memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Check for overflow (if signs are the same but result sign is different)
	if ((stackValue^memValue)&0x80000000) == 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Check for carry
	carry := (stackValue > math.MaxUint32-memValue)
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) subIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform subtraction
	result := stackValue - memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Check for overflow
	if ((stackValue^memValue)&0x80000000) != 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Check for carry/borrow
	carry := stackValue >= memValue
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) mulIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform multiplication
	result := stackValue * memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Check for overflow (simplified)
	if (memValue != 0 && result/memValue != stackValue) {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) divIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	if memValue == 0 {
		return errors.New("division by zero")
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform division
	result := stackValue / memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) addFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Convert to float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Perform addition
	resultFloat := stackFloat + memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Set flags
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) subFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Convert to float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Perform subtraction
	resultFloat := stackFloat - memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Set flags
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) mulFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Convert to float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Perform multiplication
	resultFloat := stackFloat * memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Set flags
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) divFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Convert to float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	if memFloat == 0 {
		return errors.New("division by zero")
	}

	// Perform division
	resultFloat := stackFloat / memFloat
	result := uint32(math.Float32bits(resultFloat))

	// Set flags
	c.psw.SetFlag(FLAG_FZERO, resultFloat == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, resultFloat < 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) cmpIntCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Compare values
	result := stackValue - memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	
	// Check for overflow
	if ((stackValue^memValue)&0x80000000) != 0 && ((stackValue^result)&0x80000000) != 0 {
		c.psw.SetFlag(FLAG_OVERFLOW, true)
	} else {
		c.psw.SetFlag(FLAG_OVERFLOW, false)
	}

	// Check for carry/borrow
	carry := stackValue >= memValue
	c.psw.SetFlag(FLAG_CARRY, carry)

	// Don't push result back - comparison only sets flags
	return nil
}

func (c *CPU) cmpFloatCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Convert to float32
	memFloat := math.Float32frombits(uint32(memValue))
	stackFloat := math.Float32frombits(uint32(stackValue))

	// Compare values
	result := stackFloat - memFloat

	// Set flags
	c.psw.SetFlag(FLAG_FZERO, result == 0)
	c.psw.SetFlag(FLAG_FNEGATIVE, result < 0)

	// Don't push result back - comparison only sets flags
	return nil
}

func (c *CPU) jmpCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Set IP to the new address
	c.psw.SetIP(offset)
	return nil
}

func (c *CPU) jzCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Jump if zero flag is set
	if c.psw.GetFlag(FLAG_ZERO) {
		c.psw.SetIP(offset)
	}
	return nil
}

func (c *CPU) jnzCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Jump if zero flag is not set
	if !c.psw.GetFlag(FLAG_ZERO) {
		c.psw.SetIP(offset)
	}
	return nil
}

func (c *CPU) jcCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Jump if carry flag is set
	if c.psw.GetFlag(FLAG_CARRY) {
		c.psw.SetIP(offset)
	}
	return nil
}

func (c *CPU) jncCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Jump if carry flag is not set
	if !c.psw.GetFlag(FLAG_CARRY) {
		c.psw.SetIP(offset)
	}
	return nil
}

func (c *CPU) callCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Push return address (current IP) to stack
	returnAddress := c.psw.ip
	if err := c.push(uint32(returnAddress)); err != nil {
		return err
	}

	// Set IP to the new address
	c.psw.SetIP(offset)
	return nil
}

func (c *CPU) retCommand() error {
	// Read and ignore the immediate offset (2 bytes)
	_, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Pop return address from stack
	returnAddress, err := c.pop()
	if err != nil {
		return err
	}

	// Set IP to the return address
	c.psw.SetIP(uint16(returnAddress))
	return nil
}

func (c *CPU) pushCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	value, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Push value to stack
	return c.push(value)
}

func (c *CPU) popCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Pop value from stack
	value, err := c.pop()
	if err != nil {
		return err
	}

	// Write value to memory
	return c.memory.WriteWordAt(offset, value)
}

func (c *CPU) inCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read input from user
	var input uint32
	fmt.Print("Input (integer): ")
	_, err = fmt.Scan(&input)
	if err != nil {
		return err
	}

	// Write input to memory
	return c.memory.WriteWordAt(offset, input)
}

func (c *CPU) outCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	value, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Output value
	fmt.Printf("Output: %d\n", value)
	return nil
}

func (c *CPU) andCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform AND operation
	result := stackValue & memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) orCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform OR operation
	result := stackValue | memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) xorCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read value from memory
	memValue, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform XOR operation
	result := stackValue ^ memValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) notCommand() error {
	// Read and ignore the immediate offset (2 bytes)
	_, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform NOT operation
	result := ^stackValue

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) shlCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read shift count from memory
	shiftCount, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform shift left
	result := stackValue << (shiftCount & 0x1F) // Limit shift to 31 bits

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	c.psw.SetFlag(FLAG_CARRY, (stackValue>>(32-(shiftCount&0x1F))&1) != 0)

	// Push result back to stack
	return c.push(result)
}

func (c *CPU) shrCommand() error {
	offset, err := c.readImmediateOffset()
	if err != nil {
		return err
	}

	// Read shift count from memory
	shiftCount, err := c.memory.ReadWordAt(offset)
	if err != nil {
		return err
	}

	// Pop value from stack
	stackValue, err := c.pop()
	if err != nil {
		return err
	}

	// Perform shift right
	result := stackValue >> (shiftCount & 0x1F) // Limit shift to 31 bits

	// Set flags
	c.psw.SetFlag(FLAG_ZERO, result == 0)
	c.psw.SetFlag(FLAG_NEGATIVE, (result&0x80000000) != 0)
	c.psw.SetFlag(FLAG_CARRY, (stackValue>>((shiftCount&0x1F)-1)&1) != 0)

	// Push result back to stack
	return c.push(result)
}