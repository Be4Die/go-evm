// Code generated from Assembler.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Assembler
import "github.com/antlr4-go/antlr/v4"

// BaseAssemblerListener is a complete listener for a parse tree produced by AssemblerParser.
type BaseAssemblerListener struct{}

var _ AssemblerListener = &BaseAssemblerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAssemblerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAssemblerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAssemblerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAssemblerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAssemblerListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAssemblerListener) ExitProgram(ctx *ProgramContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseAssemblerListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseAssemblerListener) ExitDirective(ctx *DirectiveContext) {}

// EnterEntryDirective is called when production entryDirective is entered.
func (s *BaseAssemblerListener) EnterEntryDirective(ctx *EntryDirectiveContext) {}

// ExitEntryDirective is called when production entryDirective is exited.
func (s *BaseAssemblerListener) ExitEntryDirective(ctx *EntryDirectiveContext) {}

// EnterSectionDirective is called when production sectionDirective is entered.
func (s *BaseAssemblerListener) EnterSectionDirective(ctx *SectionDirectiveContext) {}

// ExitSectionDirective is called when production sectionDirective is exited.
func (s *BaseAssemblerListener) ExitSectionDirective(ctx *SectionDirectiveContext) {}

// EnterDataDirective is called when production dataDirective is entered.
func (s *BaseAssemblerListener) EnterDataDirective(ctx *DataDirectiveContext) {}

// ExitDataDirective is called when production dataDirective is exited.
func (s *BaseAssemblerListener) ExitDataDirective(ctx *DataDirectiveContext) {}

// EnterEquDirective is called when production equDirective is entered.
func (s *BaseAssemblerListener) EnterEquDirective(ctx *EquDirectiveContext) {}

// ExitEquDirective is called when production equDirective is exited.
func (s *BaseAssemblerListener) ExitEquDirective(ctx *EquDirectiveContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseAssemblerListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseAssemblerListener) ExitInstruction(ctx *InstructionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseAssemblerListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseAssemblerListener) ExitLabel(ctx *LabelContext) {}

// EnterMnemonic is called when production mnemonic is entered.
func (s *BaseAssemblerListener) EnterMnemonic(ctx *MnemonicContext) {}

// ExitMnemonic is called when production mnemonic is exited.
func (s *BaseAssemblerListener) ExitMnemonic(ctx *MnemonicContext) {}

// EnterIndirectOperand is called when production indirectOperand is entered.
func (s *BaseAssemblerListener) EnterIndirectOperand(ctx *IndirectOperandContext) {}

// ExitIndirectOperand is called when production indirectOperand is exited.
func (s *BaseAssemblerListener) ExitIndirectOperand(ctx *IndirectOperandContext) {}

// EnterDirectOperand is called when production directOperand is entered.
func (s *BaseAssemblerListener) EnterDirectOperand(ctx *DirectOperandContext) {}

// ExitDirectOperand is called when production directOperand is exited.
func (s *BaseAssemblerListener) ExitDirectOperand(ctx *DirectOperandContext) {}

// EnterIndirectAddress is called when production indirectAddress is entered.
func (s *BaseAssemblerListener) EnterIndirectAddress(ctx *IndirectAddressContext) {}

// ExitIndirectAddress is called when production indirectAddress is exited.
func (s *BaseAssemblerListener) ExitIndirectAddress(ctx *IndirectAddressContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseAssemblerListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseAssemblerListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseAssemblerListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseAssemblerListener) ExitNumber(ctx *NumberContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseAssemblerListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseAssemblerListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterSectionName is called when production sectionName is entered.
func (s *BaseAssemblerListener) EnterSectionName(ctx *SectionNameContext) {}

// ExitSectionName is called when production sectionName is exited.
func (s *BaseAssemblerListener) ExitSectionName(ctx *SectionNameContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseAssemblerListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseAssemblerListener) ExitComment(ctx *CommentContext) {}

// EnterEmptyLine is called when production emptyLine is entered.
func (s *BaseAssemblerListener) EnterEmptyLine(ctx *EmptyLineContext) {}

// ExitEmptyLine is called when production emptyLine is exited.
func (s *BaseAssemblerListener) ExitEmptyLine(ctx *EmptyLineContext) {}
