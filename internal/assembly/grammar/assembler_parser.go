// Code generated from Assembler.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // Assembler
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AssemblerParser struct {
	*antlr.BaseParser
}

var AssemblerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func assemblerParserInit() {
	staticData := &AssemblerParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "','", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "ENTRY", "SECTION", "DB", "DW", "EQU", "DOT_DATA",
		"DOT_CODE", "MOV", "ADD_I", "SUB_I", "MUL_I", "DIV_I", "ADD_F", "SUB_F",
		"MUL_F", "DIV_F", "CMP_I", "CMP_F", "JMP", "JZ", "JNZ", "JC", "JNC",
		"CALL", "RET", "PUSH", "POP", "IN", "OUT", "AND", "OR", "XOR", "NOT",
		"SHL", "SHR", "HALT", "HEX_LITERAL", "BIN_LITERAL", "DEC_LITERAL", "IDENTIFIER",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS", "EOL",
	}
	staticData.RuleNames = []string{
		"program", "directive", "entryDirective", "sectionDirective", "dataDirective",
		"equDirective", "instruction", "label", "mnemonic", "operand", "indirectAddress",
		"constant", "number", "labelName", "sectionName", "comment", "emptyLine",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 150, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 40, 8, 0, 10, 0, 12,
		0, 43, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 2,
		1, 2, 3, 2, 55, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 61, 8, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 3, 4, 67, 8, 4, 1, 4, 1, 4, 3, 4, 71, 8, 4, 1, 4, 1, 4,
		3, 4, 75, 8, 4, 1, 4, 1, 4, 3, 4, 79, 8, 4, 1, 4, 1, 4, 3, 4, 83, 8, 4,
		1, 4, 5, 4, 86, 8, 4, 10, 4, 12, 4, 89, 9, 4, 1, 5, 1, 5, 3, 5, 93, 8,
		5, 1, 5, 1, 5, 3, 5, 97, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 103, 8, 6,
		1, 6, 1, 6, 3, 6, 107, 8, 6, 1, 6, 1, 6, 3, 6, 111, 8, 6, 1, 6, 3, 6, 114,
		8, 6, 1, 7, 1, 7, 3, 7, 118, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 129, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 135,
		8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 3,
		16, 146, 8, 16, 1, 16, 1, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 5, 1, 0, 7, 8, 1, 0, 12, 40, 1,
		0, 41, 43, 1, 0, 10, 11, 1, 0, 45, 46, 158, 0, 41, 1, 0, 0, 0, 2, 50, 1,
		0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 90,
		1, 0, 0, 0, 12, 106, 1, 0, 0, 0, 14, 115, 1, 0, 0, 0, 16, 121, 1, 0, 0,
		0, 18, 128, 1, 0, 0, 0, 20, 130, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24, 136,
		1, 0, 0, 0, 26, 138, 1, 0, 0, 0, 28, 140, 1, 0, 0, 0, 30, 142, 1, 0, 0,
		0, 32, 145, 1, 0, 0, 0, 34, 40, 3, 2, 1, 0, 35, 40, 3, 14, 7, 0, 36, 40,
		3, 12, 6, 0, 37, 40, 3, 30, 15, 0, 38, 40, 3, 32, 16, 0, 39, 34, 1, 0,
		0, 0, 39, 35, 1, 0, 0, 0, 39, 36, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 38,
		1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		42, 44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 5, 0, 0, 1, 45, 1, 1, 0,
		0, 0, 46, 51, 3, 4, 2, 0, 47, 51, 3, 6, 3, 0, 48, 51, 3, 8, 4, 0, 49, 51,
		3, 10, 5, 0, 50, 46, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0,
		50, 49, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 54, 5, 5, 0, 0, 53, 55, 5, 47,
		0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57,
		3, 26, 13, 0, 57, 5, 1, 0, 0, 0, 58, 60, 5, 6, 0, 0, 59, 61, 5, 47, 0,
		0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63,
		3, 28, 14, 0, 63, 7, 1, 0, 0, 0, 64, 66, 3, 26, 13, 0, 65, 67, 5, 47, 0,
		0, 66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69,
		5, 1, 0, 0, 69, 71, 1, 0, 0, 0, 70, 64, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 74, 7, 0, 0, 0, 73, 75, 5, 47, 0, 0, 74, 73, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 87, 3, 22, 11, 0,
		77, 79, 5, 47, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1,
		0, 0, 0, 80, 82, 5, 2, 0, 0, 81, 83, 5, 47, 0, 0, 82, 81, 1, 0, 0, 0, 82,
		83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 3, 22, 11, 0, 85, 78, 1, 0,
		0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 9,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92, 3, 26, 13, 0, 91, 93, 5, 47, 0,
		0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96,
		5, 9, 0, 0, 95, 97, 5, 47, 0, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 99, 3, 22, 11, 0, 99, 11, 1, 0, 0, 0, 100, 102,
		3, 26, 13, 0, 101, 103, 5, 47, 0, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 5, 1, 0, 0, 105, 107, 1, 0, 0,
		0, 106, 100, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		110, 3, 16, 8, 0, 109, 111, 5, 47, 0, 0, 110, 109, 1, 0, 0, 0, 110, 111,
		1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 114, 3, 18, 9, 0, 113, 112, 1, 0,
		0, 0, 113, 114, 1, 0, 0, 0, 114, 13, 1, 0, 0, 0, 115, 117, 3, 26, 13, 0,
		116, 118, 5, 47, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 120, 5, 1, 0, 0, 120, 15, 1, 0, 0, 0, 121, 122, 7,
		1, 0, 0, 122, 17, 1, 0, 0, 0, 123, 124, 5, 3, 0, 0, 124, 125, 3, 20, 10,
		0, 125, 126, 5, 4, 0, 0, 126, 129, 1, 0, 0, 0, 127, 129, 3, 22, 11, 0,
		128, 123, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 19, 1, 0, 0, 0, 130, 131,
		3, 22, 11, 0, 131, 21, 1, 0, 0, 0, 132, 135, 3, 24, 12, 0, 133, 135, 3,
		26, 13, 0, 134, 132, 1, 0, 0, 0, 134, 133, 1, 0, 0, 0, 135, 23, 1, 0, 0,
		0, 136, 137, 7, 2, 0, 0, 137, 25, 1, 0, 0, 0, 138, 139, 5, 44, 0, 0, 139,
		27, 1, 0, 0, 0, 140, 141, 7, 3, 0, 0, 141, 29, 1, 0, 0, 0, 142, 143, 7,
		4, 0, 0, 143, 31, 1, 0, 0, 0, 144, 146, 5, 47, 0, 0, 145, 144, 1, 0, 0,
		0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 48, 0, 0, 148,
		33, 1, 0, 0, 0, 21, 39, 41, 50, 54, 60, 66, 70, 74, 78, 82, 87, 92, 96,
		102, 106, 110, 113, 117, 128, 134, 145,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AssemblerParserInit initializes any static state used to implement AssemblerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAssemblerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AssemblerParserInit() {
	staticData := &AssemblerParserStaticData
	staticData.once.Do(assemblerParserInit)
}

// NewAssemblerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAssemblerParser(input antlr.TokenStream) *AssemblerParser {
	AssemblerParserInit()
	this := new(AssemblerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AssemblerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Assembler.g4"

	return this
}

// AssemblerParser tokens.
const (
	AssemblerParserEOF           = antlr.TokenEOF
	AssemblerParserT__0          = 1
	AssemblerParserT__1          = 2
	AssemblerParserT__2          = 3
	AssemblerParserT__3          = 4
	AssemblerParserENTRY         = 5
	AssemblerParserSECTION       = 6
	AssemblerParserDB            = 7
	AssemblerParserDW            = 8
	AssemblerParserEQU           = 9
	AssemblerParserDOT_DATA      = 10
	AssemblerParserDOT_CODE      = 11
	AssemblerParserMOV           = 12
	AssemblerParserADD_I         = 13
	AssemblerParserSUB_I         = 14
	AssemblerParserMUL_I         = 15
	AssemblerParserDIV_I         = 16
	AssemblerParserADD_F         = 17
	AssemblerParserSUB_F         = 18
	AssemblerParserMUL_F         = 19
	AssemblerParserDIV_F         = 20
	AssemblerParserCMP_I         = 21
	AssemblerParserCMP_F         = 22
	AssemblerParserJMP           = 23
	AssemblerParserJZ            = 24
	AssemblerParserJNZ           = 25
	AssemblerParserJC            = 26
	AssemblerParserJNC           = 27
	AssemblerParserCALL          = 28
	AssemblerParserRET           = 29
	AssemblerParserPUSH          = 30
	AssemblerParserPOP           = 31
	AssemblerParserIN            = 32
	AssemblerParserOUT           = 33
	AssemblerParserAND           = 34
	AssemblerParserOR            = 35
	AssemblerParserXOR           = 36
	AssemblerParserNOT           = 37
	AssemblerParserSHL           = 38
	AssemblerParserSHR           = 39
	AssemblerParserHALT          = 40
	AssemblerParserHEX_LITERAL   = 41
	AssemblerParserBIN_LITERAL   = 42
	AssemblerParserDEC_LITERAL   = 43
	AssemblerParserIDENTIFIER    = 44
	AssemblerParserLINE_COMMENT  = 45
	AssemblerParserBLOCK_COMMENT = 46
	AssemblerParserWS            = 47
	AssemblerParserEOL           = 48
)

// AssemblerParser rules.
const (
	AssemblerParserRULE_program          = 0
	AssemblerParserRULE_directive        = 1
	AssemblerParserRULE_entryDirective   = 2
	AssemblerParserRULE_sectionDirective = 3
	AssemblerParserRULE_dataDirective    = 4
	AssemblerParserRULE_equDirective     = 5
	AssemblerParserRULE_instruction      = 6
	AssemblerParserRULE_label            = 7
	AssemblerParserRULE_mnemonic         = 8
	AssemblerParserRULE_operand          = 9
	AssemblerParserRULE_indirectAddress  = 10
	AssemblerParserRULE_constant         = 11
	AssemblerParserRULE_number           = 12
	AssemblerParserRULE_labelName        = 13
	AssemblerParserRULE_sectionName      = 14
	AssemblerParserRULE_comment          = 15
	AssemblerParserRULE_emptyLine        = 16
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDirective() []IDirectiveContext
	Directive(i int) IDirectiveContext
	AllLabel() []ILabelContext
	Label(i int) ILabelContext
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext
	AllComment() []ICommentContext
	Comment(i int) ICommentContext
	AllEmptyLine() []IEmptyLineContext
	EmptyLine(i int) IEmptyLineContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(AssemblerParserEOF, 0)
}

func (s *ProgramContext) AllDirective() []IDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirectiveContext); ok {
			tst[i] = t.(IDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Directive(i int) IDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *ProgramContext) AllLabel() []ILabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelContext); ok {
			len++
		}
	}

	tst := make([]ILabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelContext); ok {
			tst[i] = t.(ILabelContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Label(i int) ILabelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ProgramContext) AllEmptyLine() []IEmptyLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyLineContext); ok {
			len++
		}
	}

	tst := make([]IEmptyLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyLineContext); ok {
			tst[i] = t.(IEmptyLineContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) EmptyLine(i int) IEmptyLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyLineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *AssemblerParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AssemblerParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&547556790628832) != 0 {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(34)
				p.Directive()
			}

		case 2:
			{
				p.SetState(35)
				p.Label()
			}

		case 3:
			{
				p.SetState(36)
				p.Instruction()
			}

		case 4:
			{
				p.SetState(37)
				p.Comment()
			}

		case 5:
			{
				p.SetState(38)
				p.EmptyLine()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(AssemblerParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EntryDirective() IEntryDirectiveContext
	SectionDirective() ISectionDirectiveContext
	DataDirective() IDataDirectiveContext
	EquDirective() IEquDirectiveContext

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_directive
	return p
}

func InitEmptyDirectiveContext(p *DirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_directive
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) EntryDirective() IEntryDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntryDirectiveContext)
}

func (s *DirectiveContext) SectionDirective() ISectionDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionDirectiveContext)
}

func (s *DirectiveContext) DataDirective() IDataDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDirectiveContext)
}

func (s *DirectiveContext) EquDirective() IEquDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEquDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEquDirectiveContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *AssemblerParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AssemblerParserRULE_directive)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.EntryDirective()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.SectionDirective()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.DataDirective()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.EquDirective()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntryDirectiveContext is an interface to support dynamic dispatch.
type IEntryDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTRY() antlr.TerminalNode
	LabelName() ILabelNameContext
	WS() antlr.TerminalNode

	// IsEntryDirectiveContext differentiates from other interfaces.
	IsEntryDirectiveContext()
}

type EntryDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryDirectiveContext() *EntryDirectiveContext {
	var p = new(EntryDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_entryDirective
	return p
}

func InitEmptyEntryDirectiveContext(p *EntryDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_entryDirective
}

func (*EntryDirectiveContext) IsEntryDirectiveContext() {}

func NewEntryDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryDirectiveContext {
	var p = new(EntryDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_entryDirective

	return p
}

func (s *EntryDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryDirectiveContext) ENTRY() antlr.TerminalNode {
	return s.GetToken(AssemblerParserENTRY, 0)
}

func (s *EntryDirectiveContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *EntryDirectiveContext) WS() antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, 0)
}

func (s *EntryDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterEntryDirective(s)
	}
}

func (s *EntryDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitEntryDirective(s)
	}
}

func (p *AssemblerParser) EntryDirective() (localctx IEntryDirectiveContext) {
	localctx = NewEntryDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AssemblerParserRULE_entryDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(AssemblerParserENTRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(53)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(56)
		p.LabelName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISectionDirectiveContext is an interface to support dynamic dispatch.
type ISectionDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SECTION() antlr.TerminalNode
	SectionName() ISectionNameContext
	WS() antlr.TerminalNode

	// IsSectionDirectiveContext differentiates from other interfaces.
	IsSectionDirectiveContext()
}

type SectionDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionDirectiveContext() *SectionDirectiveContext {
	var p = new(SectionDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_sectionDirective
	return p
}

func InitEmptySectionDirectiveContext(p *SectionDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_sectionDirective
}

func (*SectionDirectiveContext) IsSectionDirectiveContext() {}

func NewSectionDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionDirectiveContext {
	var p = new(SectionDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_sectionDirective

	return p
}

func (s *SectionDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionDirectiveContext) SECTION() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSECTION, 0)
}

func (s *SectionDirectiveContext) SectionName() ISectionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionNameContext)
}

func (s *SectionDirectiveContext) WS() antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, 0)
}

func (s *SectionDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterSectionDirective(s)
	}
}

func (s *SectionDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitSectionDirective(s)
	}
}

func (p *AssemblerParser) SectionDirective() (localctx ISectionDirectiveContext) {
	localctx = NewSectionDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AssemblerParserRULE_sectionDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(AssemblerParserSECTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(59)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(62)
		p.SectionName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataDirectiveContext is an interface to support dynamic dispatch.
type IDataDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	DB() antlr.TerminalNode
	DW() antlr.TerminalNode
	LabelName() ILabelNameContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsDataDirectiveContext differentiates from other interfaces.
	IsDataDirectiveContext()
}

type DataDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataDirectiveContext() *DataDirectiveContext {
	var p = new(DataDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_dataDirective
	return p
}

func InitEmptyDataDirectiveContext(p *DataDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_dataDirective
}

func (*DataDirectiveContext) IsDataDirectiveContext() {}

func NewDataDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataDirectiveContext {
	var p = new(DataDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_dataDirective

	return p
}

func (s *DataDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DataDirectiveContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *DataDirectiveContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DataDirectiveContext) DB() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDB, 0)
}

func (s *DataDirectiveContext) DW() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDW, 0)
}

func (s *DataDirectiveContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *DataDirectiveContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserWS)
}

func (s *DataDirectiveContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, i)
}

func (s *DataDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterDataDirective(s)
	}
}

func (s *DataDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitDataDirective(s)
	}
}

func (p *AssemblerParser) DataDirective() (localctx IDataDirectiveContext) {
	localctx = NewDataDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AssemblerParserRULE_dataDirective)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserIDENTIFIER {
		{
			p.SetState(64)
			p.LabelName()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AssemblerParserWS {
			{
				p.SetState(65)
				p.Match(AssemblerParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(68)
			p.Match(AssemblerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(72)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblerParserDB || _la == AssemblerParserDW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(73)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(76)
		p.Constant()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == AssemblerParserWS {
				{
					p.SetState(77)
					p.Match(AssemblerParserWS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(80)
				p.Match(AssemblerParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == AssemblerParserWS {
				{
					p.SetState(81)
					p.Match(AssemblerParserWS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(84)
				p.Constant()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEquDirectiveContext is an interface to support dynamic dispatch.
type IEquDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabelName() ILabelNameContext
	EQU() antlr.TerminalNode
	Constant() IConstantContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsEquDirectiveContext differentiates from other interfaces.
	IsEquDirectiveContext()
}

type EquDirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquDirectiveContext() *EquDirectiveContext {
	var p = new(EquDirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_equDirective
	return p
}

func InitEmptyEquDirectiveContext(p *EquDirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_equDirective
}

func (*EquDirectiveContext) IsEquDirectiveContext() {}

func NewEquDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquDirectiveContext {
	var p = new(EquDirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_equDirective

	return p
}

func (s *EquDirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *EquDirectiveContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *EquDirectiveContext) EQU() antlr.TerminalNode {
	return s.GetToken(AssemblerParserEQU, 0)
}

func (s *EquDirectiveContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EquDirectiveContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserWS)
}

func (s *EquDirectiveContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, i)
}

func (s *EquDirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquDirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquDirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterEquDirective(s)
	}
}

func (s *EquDirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitEquDirective(s)
	}
}

func (p *AssemblerParser) EquDirective() (localctx IEquDirectiveContext) {
	localctx = NewEquDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AssemblerParserRULE_equDirective)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.LabelName()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(91)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(94)
		p.Match(AssemblerParserEQU)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(95)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(98)
		p.Constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mnemonic() IMnemonicContext
	LabelName() ILabelNameContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode
	Operand() IOperandContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Mnemonic() IMnemonicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMnemonicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMnemonicContext)
}

func (s *InstructionContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *InstructionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(AssemblerParserWS)
}

func (s *InstructionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, i)
}

func (s *InstructionContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *AssemblerParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AssemblerParserRULE_instruction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserIDENTIFIER {
		{
			p.SetState(100)
			p.LabelName()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AssemblerParserWS {
			{
				p.SetState(101)
				p.Match(AssemblerParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(104)
			p.Match(AssemblerParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(108)
		p.Mnemonic()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(109)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(112)
			p.Operand()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabelName() ILabelNameContext
	WS() antlr.TerminalNode

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *LabelContext) WS() antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *AssemblerParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AssemblerParserRULE_label)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.LabelName()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(116)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(119)
		p.Match(AssemblerParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMnemonicContext is an interface to support dynamic dispatch.
type IMnemonicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MOV() antlr.TerminalNode
	ADD_I() antlr.TerminalNode
	SUB_I() antlr.TerminalNode
	MUL_I() antlr.TerminalNode
	DIV_I() antlr.TerminalNode
	ADD_F() antlr.TerminalNode
	SUB_F() antlr.TerminalNode
	MUL_F() antlr.TerminalNode
	DIV_F() antlr.TerminalNode
	CMP_I() antlr.TerminalNode
	CMP_F() antlr.TerminalNode
	JMP() antlr.TerminalNode
	JZ() antlr.TerminalNode
	JNZ() antlr.TerminalNode
	JC() antlr.TerminalNode
	JNC() antlr.TerminalNode
	CALL() antlr.TerminalNode
	RET() antlr.TerminalNode
	PUSH() antlr.TerminalNode
	POP() antlr.TerminalNode
	IN() antlr.TerminalNode
	OUT() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	XOR() antlr.TerminalNode
	NOT() antlr.TerminalNode
	SHL() antlr.TerminalNode
	SHR() antlr.TerminalNode
	HALT() antlr.TerminalNode

	// IsMnemonicContext differentiates from other interfaces.
	IsMnemonicContext()
}

type MnemonicContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMnemonicContext() *MnemonicContext {
	var p = new(MnemonicContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_mnemonic
	return p
}

func InitEmptyMnemonicContext(p *MnemonicContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_mnemonic
}

func (*MnemonicContext) IsMnemonicContext() {}

func NewMnemonicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MnemonicContext {
	var p = new(MnemonicContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_mnemonic

	return p
}

func (s *MnemonicContext) GetParser() antlr.Parser { return s.parser }

func (s *MnemonicContext) MOV() antlr.TerminalNode {
	return s.GetToken(AssemblerParserMOV, 0)
}

func (s *MnemonicContext) ADD_I() antlr.TerminalNode {
	return s.GetToken(AssemblerParserADD_I, 0)
}

func (s *MnemonicContext) SUB_I() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSUB_I, 0)
}

func (s *MnemonicContext) MUL_I() antlr.TerminalNode {
	return s.GetToken(AssemblerParserMUL_I, 0)
}

func (s *MnemonicContext) DIV_I() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDIV_I, 0)
}

func (s *MnemonicContext) ADD_F() antlr.TerminalNode {
	return s.GetToken(AssemblerParserADD_F, 0)
}

func (s *MnemonicContext) SUB_F() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSUB_F, 0)
}

func (s *MnemonicContext) MUL_F() antlr.TerminalNode {
	return s.GetToken(AssemblerParserMUL_F, 0)
}

func (s *MnemonicContext) DIV_F() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDIV_F, 0)
}

func (s *MnemonicContext) CMP_I() antlr.TerminalNode {
	return s.GetToken(AssemblerParserCMP_I, 0)
}

func (s *MnemonicContext) CMP_F() antlr.TerminalNode {
	return s.GetToken(AssemblerParserCMP_F, 0)
}

func (s *MnemonicContext) JMP() antlr.TerminalNode {
	return s.GetToken(AssemblerParserJMP, 0)
}

func (s *MnemonicContext) JZ() antlr.TerminalNode {
	return s.GetToken(AssemblerParserJZ, 0)
}

func (s *MnemonicContext) JNZ() antlr.TerminalNode {
	return s.GetToken(AssemblerParserJNZ, 0)
}

func (s *MnemonicContext) JC() antlr.TerminalNode {
	return s.GetToken(AssemblerParserJC, 0)
}

func (s *MnemonicContext) JNC() antlr.TerminalNode {
	return s.GetToken(AssemblerParserJNC, 0)
}

func (s *MnemonicContext) CALL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserCALL, 0)
}

func (s *MnemonicContext) RET() antlr.TerminalNode {
	return s.GetToken(AssemblerParserRET, 0)
}

func (s *MnemonicContext) PUSH() antlr.TerminalNode {
	return s.GetToken(AssemblerParserPUSH, 0)
}

func (s *MnemonicContext) POP() antlr.TerminalNode {
	return s.GetToken(AssemblerParserPOP, 0)
}

func (s *MnemonicContext) IN() antlr.TerminalNode {
	return s.GetToken(AssemblerParserIN, 0)
}

func (s *MnemonicContext) OUT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserOUT, 0)
}

func (s *MnemonicContext) AND() antlr.TerminalNode {
	return s.GetToken(AssemblerParserAND, 0)
}

func (s *MnemonicContext) OR() antlr.TerminalNode {
	return s.GetToken(AssemblerParserOR, 0)
}

func (s *MnemonicContext) XOR() antlr.TerminalNode {
	return s.GetToken(AssemblerParserXOR, 0)
}

func (s *MnemonicContext) NOT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserNOT, 0)
}

func (s *MnemonicContext) SHL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSHL, 0)
}

func (s *MnemonicContext) SHR() antlr.TerminalNode {
	return s.GetToken(AssemblerParserSHR, 0)
}

func (s *MnemonicContext) HALT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserHALT, 0)
}

func (s *MnemonicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MnemonicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MnemonicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterMnemonic(s)
	}
}

func (s *MnemonicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitMnemonic(s)
	}
}

func (p *AssemblerParser) Mnemonic() (localctx IMnemonicContext) {
	localctx = NewMnemonicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AssemblerParserRULE_mnemonic)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023251456) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DirectOperandContext struct {
	OperandContext
}

func NewDirectOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DirectOperandContext {
	var p = new(DirectOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *DirectOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectOperandContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DirectOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterDirectOperand(s)
	}
}

func (s *DirectOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitDirectOperand(s)
	}
}

type IndirectOperandContext struct {
	OperandContext
}

func NewIndirectOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndirectOperandContext {
	var p = new(IndirectOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *IndirectOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndirectOperandContext) IndirectAddress() IIndirectAddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndirectAddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndirectAddressContext)
}

func (s *IndirectOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterIndirectOperand(s)
	}
}

func (s *IndirectOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitIndirectOperand(s)
	}
}

func (p *AssemblerParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AssemblerParserRULE_operand)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AssemblerParserT__2:
		localctx = NewIndirectOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(AssemblerParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.IndirectAddress()
		}
		{
			p.SetState(125)
			p.Match(AssemblerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AssemblerParserHEX_LITERAL, AssemblerParserBIN_LITERAL, AssemblerParserDEC_LITERAL, AssemblerParserIDENTIFIER:
		localctx = NewDirectOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndirectAddressContext is an interface to support dynamic dispatch.
type IIndirectAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext

	// IsIndirectAddressContext differentiates from other interfaces.
	IsIndirectAddressContext()
}

type IndirectAddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndirectAddressContext() *IndirectAddressContext {
	var p = new(IndirectAddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_indirectAddress
	return p
}

func InitEmptyIndirectAddressContext(p *IndirectAddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_indirectAddress
}

func (*IndirectAddressContext) IsIndirectAddressContext() {}

func NewIndirectAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndirectAddressContext {
	var p = new(IndirectAddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_indirectAddress

	return p
}

func (s *IndirectAddressContext) GetParser() antlr.Parser { return s.parser }

func (s *IndirectAddressContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *IndirectAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndirectAddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndirectAddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterIndirectAddress(s)
	}
}

func (s *IndirectAddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitIndirectAddress(s)
	}
}

func (p *AssemblerParser) IndirectAddress() (localctx IIndirectAddressContext) {
	localctx = NewIndirectAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AssemblerParserRULE_indirectAddress)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Constant()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	LabelName() ILabelNameContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ConstantContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *AssemblerParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AssemblerParserRULE_constant)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AssemblerParserHEX_LITERAL, AssemblerParserBIN_LITERAL, AssemblerParserDEC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Number()
		}

	case AssemblerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.LabelName()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HEX_LITERAL() antlr.TerminalNode
	BIN_LITERAL() antlr.TerminalNode
	DEC_LITERAL() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserHEX_LITERAL, 0)
}

func (s *NumberContext) BIN_LITERAL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserBIN_LITERAL, 0)
}

func (s *NumberContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDEC_LITERAL, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *AssemblerParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AssemblerParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15393162788864) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelNameContext is an interface to support dynamic dispatch.
type ILabelNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLabelNameContext differentiates from other interfaces.
	IsLabelNameContext()
}

type LabelNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelNameContext() *LabelNameContext {
	var p = new(LabelNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_labelName
	return p
}

func InitEmptyLabelNameContext(p *LabelNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_labelName
}

func (*LabelNameContext) IsLabelNameContext() {}

func NewLabelNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameContext {
	var p = new(LabelNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_labelName

	return p
}

func (s *LabelNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AssemblerParserIDENTIFIER, 0)
}

func (s *LabelNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterLabelName(s)
	}
}

func (s *LabelNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitLabelName(s)
	}
}

func (p *AssemblerParser) LabelName() (localctx ILabelNameContext) {
	localctx = NewLabelNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AssemblerParserRULE_labelName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(AssemblerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISectionNameContext is an interface to support dynamic dispatch.
type ISectionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT_DATA() antlr.TerminalNode
	DOT_CODE() antlr.TerminalNode

	// IsSectionNameContext differentiates from other interfaces.
	IsSectionNameContext()
}

type SectionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionNameContext() *SectionNameContext {
	var p = new(SectionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_sectionName
	return p
}

func InitEmptySectionNameContext(p *SectionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_sectionName
}

func (*SectionNameContext) IsSectionNameContext() {}

func NewSectionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionNameContext {
	var p = new(SectionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_sectionName

	return p
}

func (s *SectionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionNameContext) DOT_DATA() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDOT_DATA, 0)
}

func (s *SectionNameContext) DOT_CODE() antlr.TerminalNode {
	return s.GetToken(AssemblerParserDOT_CODE, 0)
}

func (s *SectionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterSectionName(s)
	}
}

func (s *SectionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitSectionName(s)
	}
}

func (p *AssemblerParser) SectionName() (localctx ISectionNameContext) {
	localctx = NewSectionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AssemblerParserRULE_sectionName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblerParserDOT_DATA || _la == AssemblerParserDOT_CODE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LINE_COMMENT() antlr.TerminalNode
	BLOCK_COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserLINE_COMMENT, 0)
}

func (s *CommentContext) BLOCK_COMMENT() antlr.TerminalNode {
	return s.GetToken(AssemblerParserBLOCK_COMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *AssemblerParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AssemblerParserRULE_comment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AssemblerParserLINE_COMMENT || _la == AssemblerParserBLOCK_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyLineContext is an interface to support dynamic dispatch.
type IEmptyLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOL() antlr.TerminalNode
	WS() antlr.TerminalNode

	// IsEmptyLineContext differentiates from other interfaces.
	IsEmptyLineContext()
}

type EmptyLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyLineContext() *EmptyLineContext {
	var p = new(EmptyLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_emptyLine
	return p
}

func InitEmptyEmptyLineContext(p *EmptyLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AssemblerParserRULE_emptyLine
}

func (*EmptyLineContext) IsEmptyLineContext() {}

func NewEmptyLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyLineContext {
	var p = new(EmptyLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AssemblerParserRULE_emptyLine

	return p
}

func (s *EmptyLineContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyLineContext) EOL() antlr.TerminalNode {
	return s.GetToken(AssemblerParserEOL, 0)
}

func (s *EmptyLineContext) WS() antlr.TerminalNode {
	return s.GetToken(AssemblerParserWS, 0)
}

func (s *EmptyLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.EnterEmptyLine(s)
	}
}

func (s *EmptyLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AssemblerListener); ok {
		listenerT.ExitEmptyLine(s)
	}
}

func (p *AssemblerParser) EmptyLine() (localctx IEmptyLineContext) {
	localctx = NewEmptyLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AssemblerParserRULE_emptyLine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AssemblerParserWS {
		{
			p.SetState(144)
			p.Match(AssemblerParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(147)
		p.Match(AssemblerParserEOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
