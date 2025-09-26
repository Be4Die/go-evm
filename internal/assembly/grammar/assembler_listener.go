// Code generated from Assembler.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Assembler
import "github.com/antlr4-go/antlr/v4"

// AssemblerListener is a complete listener for a parse tree produced by AssemblerParser.
type AssemblerListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterEntryDirective is called when entering the entryDirective production.
	EnterEntryDirective(c *EntryDirectiveContext)

	// EnterSectionDirective is called when entering the sectionDirective production.
	EnterSectionDirective(c *SectionDirectiveContext)

	// EnterDataDirective is called when entering the dataDirective production.
	EnterDataDirective(c *DataDirectiveContext)

	// EnterEquDirective is called when entering the equDirective production.
	EnterEquDirective(c *EquDirectiveContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterMnemonic is called when entering the mnemonic production.
	EnterMnemonic(c *MnemonicContext)

	// EnterIndirectOperand is called when entering the indirectOperand production.
	EnterIndirectOperand(c *IndirectOperandContext)

	// EnterDirectOperand is called when entering the directOperand production.
	EnterDirectOperand(c *DirectOperandContext)

	// EnterIndirectAddress is called when entering the indirectAddress production.
	EnterIndirectAddress(c *IndirectAddressContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterSectionName is called when entering the sectionName production.
	EnterSectionName(c *SectionNameContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterEmptyLine is called when entering the emptyLine production.
	EnterEmptyLine(c *EmptyLineContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitEntryDirective is called when exiting the entryDirective production.
	ExitEntryDirective(c *EntryDirectiveContext)

	// ExitSectionDirective is called when exiting the sectionDirective production.
	ExitSectionDirective(c *SectionDirectiveContext)

	// ExitDataDirective is called when exiting the dataDirective production.
	ExitDataDirective(c *DataDirectiveContext)

	// ExitEquDirective is called when exiting the equDirective production.
	ExitEquDirective(c *EquDirectiveContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitMnemonic is called when exiting the mnemonic production.
	ExitMnemonic(c *MnemonicContext)

	// ExitIndirectOperand is called when exiting the indirectOperand production.
	ExitIndirectOperand(c *IndirectOperandContext)

	// ExitDirectOperand is called when exiting the directOperand production.
	ExitDirectOperand(c *DirectOperandContext)

	// ExitIndirectAddress is called when exiting the indirectAddress production.
	ExitIndirectAddress(c *IndirectAddressContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitSectionName is called when exiting the sectionName production.
	ExitSectionName(c *SectionNameContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitEmptyLine is called when exiting the emptyLine production.
	ExitEmptyLine(c *EmptyLineContext)
}
