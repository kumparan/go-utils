package utils

import (
	"encoding/json"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	_commaSeparator    = ","
	_dotSeparator      = "."
	_newline           = "\n"
	_punctuationMarks  = ".,:;!?"
	_sentenceSeparator = ". "
	_spaceSeparator    = " "

	_nodeTypeHeadingLarge  = "heading-large"
	_nodeTypeHeadingMedium = "heading-medium"
	_nodeTypeParagraph     = "paragraph"
	_nodeTypeBulletedList  = "bulleted-list"
	_nodeTypeNumberedList  = "numbered-list"
	_nodeTypeListItem      = "list-item"
	_nodeTypeInline        = "inline"
	_nodeTypeCaption       = "caption"
	_nodeTypeFigure        = "figure"
)

// SlateLeaf represents a text element with optional formatting.
type SlateLeaf struct {
	Object string `json:"object"`
	Text   string `json:"text"`
	Marks  []any  `json:"marks"`
}

// SlateNode represents a hierarchical document structure.
type SlateNode struct {
	Object     string      `json:"object"`
	Type       string      `json:"type"`
	SlateNodes []SlateNode `json:"nodes,omitempty"`
	Leaves     []SlateLeaf `json:"leaves,omitempty"`

	isLastListElement bool
}

// SlateDocument represents the root structure of a Slate document.
type SlateDocument struct {
	Document struct {
		Nodes []SlateNode `json:"nodes"`
	} `json:"document"`
}

// ParseSlateDocument parses a Slate document JSON string into a SlateDocument struct.
func ParseSlateDocument(documentJSON string) (*SlateDocument, error) {
	var wrappedDocument SlateDocument
	if err := json.Unmarshal([]byte(documentJSON), &wrappedDocument); err != nil {
		log.WithField("documentJSON", documentJSON).Error(err)
		return nil, err
	}

	return &wrappedDocument, nil
}

// ToPlainText processes and formats a Slate document JSON string.
func (slateDocument *SlateDocument) ToPlainText() (string, error) {
	text := serializeSlateNodes(slateDocument.Document.Nodes, _newline, _spaceSeparator)
	text = regexp.MustCompile(`\n+`).ReplaceAllString(text, _newline)

	return strings.TrimSpace(text), nil
}

// serializeSlateNodes recursively processes nodes and its content into a plain-text format.
func serializeSlateNodes(nodes []SlateNode, nodeSeparator, leaveSeparator string) string {
	var result strings.Builder
	modifiedSlateNodeSeparator := nodeSeparator

	for _, node := range nodes {
		// Handle paragraph nodes by ensuring they end with punctuation.
		if node.Type == _nodeTypeParagraph {
			ensureEndsWithPunctuation(&node)
		}

		// Recursively process child nodes.
		switch {
		case len(node.SlateNodes) > 0:
			switch node.Type {
			case _nodeTypeHeadingLarge, _nodeTypeCaption, _nodeTypeFigure:
				continue
			case _nodeTypeHeadingMedium:
				modifiedSlateNodeSeparator = _sentenceSeparator
			case _nodeTypeBulletedList, _nodeTypeNumberedList:
				node.SlateNodes[len(node.SlateNodes)-1].isLastListElement = true
				temp := serializeSlateNodes(node.SlateNodes, modifiedSlateNodeSeparator, leaveSeparator)
				cleaned := cleanUpList(temp)
				result.WriteString(cleaned)
				continue
			case _nodeTypeInline:
				modifiedSlateNodeSeparator = _spaceSeparator
				result.WriteString(serializeSlateNodes(node.SlateNodes, modifiedSlateNodeSeparator, leaveSeparator) + leaveSeparator)
			case _nodeTypeListItem:
				modifiedSlateNodeSeparator = _commaSeparator
				if node.isLastListElement {
					modifiedSlateNodeSeparator = _sentenceSeparator
				}
				result.WriteString(serializeSlateNodes(node.SlateNodes, modifiedSlateNodeSeparator, leaveSeparator) + leaveSeparator)
			default:
				res := serializeSlateNodes(node.SlateNodes, _spaceSeparator, leaveSeparator)
				res = strings.TrimSpace(res)
				if modifiedSlateNodeSeparator == _commaSeparator && endsWithPunctuation(res) {
					res = res[:len(res)-1] + modifiedSlateNodeSeparator
				} else if modifiedSlateNodeSeparator == _newline {
					res += modifiedSlateNodeSeparator
				}
				result.WriteString(res)
			}
		case len(node.Leaves) > 0:
			result.WriteString(serializeLeaves(node.Leaves, leaveSeparator))
		}
	}

	return result.String()
}

// serializeLeaves joins leaf texts with the specified separator.
func serializeLeaves(leaves []SlateLeaf, separator string) string {
	var result []string
	for _, leaf := range leaves {
		text := leaf.Text
		if text != "" {
			result = append(result, text)
		}
	}

	return strings.Join(result, separator)
}

// ensureEndsWithPunctuation ensures that the last leaf of a paragraph node ends with punctuation.
func ensureEndsWithPunctuation(node *SlateNode) {
	if len(node.Leaves) > 0 {
		lastSlateLeaf := node.Leaves[len(node.Leaves)-1]
		if text := lastSlateLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastSlateLeaf.Text += _sentenceSeparator
			}
		}
		node.Leaves[len(node.Leaves)-1] = lastSlateLeaf
	}

	if len(node.SlateNodes) > 0 && len(node.SlateNodes[len(node.SlateNodes)-1].Leaves) > 0 {
		lastSlateLeaf := node.SlateNodes[len(node.SlateNodes)-1].Leaves[len(node.SlateNodes[len(node.SlateNodes)-1].Leaves)-1]
		if text := lastSlateLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastSlateLeaf.Text += _sentenceSeparator
			}
		}
		node.SlateNodes[len(node.SlateNodes)-1].Leaves[len(node.SlateNodes[len(node.SlateNodes)-1].Leaves)-1] = lastSlateLeaf
	}
}

// cleans up the serialized list by removing excessive punctuation.
func cleanUpList(text string) string {
	cleaned := regexp.MustCompile(`\.+`).ReplaceAllString(text, _sentenceSeparator)
	return regexp.MustCompile(`\.\s`).ReplaceAllString(cleaned, _dotSeparator)
}

func endsWithPunctuation(text string) bool {
	if len(text) == 0 {
		return false
	}

	lastChar := string(text[len(text)-1])
	return strings.Contains(_punctuationMarks, lastChar)
}
