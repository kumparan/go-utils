package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	_multipleDotsRegex = regexp.MustCompile(`\.+`)
	_dotSpaceRegex     = regexp.MustCompile(`\.\s`)
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
	_nodeTypeLink          = "link"
)

// SlateMark represents a text decoration.
type SlateMark struct {
	Type string `json:"type"`
}

// SlateLeaf represents a text element with optional formatting.
type SlateLeaf struct {
	Object string      `json:"object"`
	Text   string      `json:"text"`
	Marks  []SlateMark `json:"marks"`
}

// SlateNode represents a hierarchical document structure.
type SlateNode struct {
	Object string         `json:"object"`
	Type   string         `json:"type"`
	Data   map[string]any `json:"data,omitempty"`
	Nodes  []SlateNode    `json:"nodes,omitempty"`
	Leaves []SlateLeaf    `json:"leaves,omitempty"`

	isLastInList bool
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

// ToPlainText converts a Slate document into a plain-text format.
func (slateDocument SlateDocument) ToPlainText() (string, error) {
	text := serializeSlateNodes(slateDocument.Document.Nodes, _newline, _spaceSeparator)
	text = regexp.MustCompile(`\n+`).ReplaceAllString(text, _newline)

	return strings.TrimSpace(text), nil
}

// ToMarkdown converts a Slate document into a markdown format.
func (slateDocument SlateDocument) ToMarkdown() (string, error) {
	text := serializeSlateNodesMarkdown(slateDocument.Document.Nodes)
	// Replace excessive newlines
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")

	return strings.TrimSpace(text), nil
}

// serializeSlateNodesMarkdown recursively processes nodes into a markdown format.
func serializeSlateNodesMarkdown(nodes []SlateNode) string {
	var result strings.Builder
	for _, node := range nodes {
		switch node.Type {
		case _nodeTypeHeadingLarge:
			result.WriteString("# " + serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves) + "\n\n")
		case _nodeTypeHeadingMedium:
			result.WriteString("## " + serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves) + "\n\n")
		case _nodeTypeParagraph:
			result.WriteString(serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves) + "\n\n")
		case _nodeTypeBulletedList:
			for _, item := range node.Nodes {
				if item.Type == _nodeTypeListItem {
					result.WriteString("* " + serializeSlateNodesMarkdown(item.Nodes) + serializeSlateLeavesMarkdown(item.Leaves) + "\n")
				}
			}
			result.WriteString("\n")
		case _nodeTypeNumberedList:
			for i, item := range node.Nodes {
				if item.Type == _nodeTypeListItem {
					fmt.Fprintf(&result, "%d. %s\n", i+1, serializeSlateNodesMarkdown(item.Nodes)+serializeSlateLeavesMarkdown(item.Leaves))
				}
			}
			result.WriteString("\n")
		case _nodeTypeLink:
			href := ""
			if node.Data != nil {
				if h, ok := node.Data["href"].(string); ok {
					href = h
				}
			}
			content := serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves)
			result.WriteString("[" + content + "](" + href + ")")
		case _nodeTypeInline:
			result.WriteString(serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves))
		default:
			// For other types or when it's just a container
			if node.Object == "text" {
				result.WriteString(serializeSlateLeavesMarkdown(node.Leaves))
			} else {
				content := serializeSlateNodesMarkdown(node.Nodes) + serializeSlateLeavesMarkdown(node.Leaves)
				if content != "" {
					result.WriteString(content)
				}
			}
		}
	}
	return result.String()
}

// serializeSlateLeavesMarkdown joins leaf texts with markdown marks.
func serializeSlateLeavesMarkdown(leaves []SlateLeaf) string {
	var result strings.Builder
	for _, leaf := range leaves {
		text := leaf.Text
		if text == "" {
			continue
		}

		// Apply marks in a specific order to ensure consistent nesting
		for _, mark := range leaf.Marks {
			switch mark.Type {
			case "bold":
				text = "**" + text + "**"
			case "italic":
				text = "*" + text + "*"
			case "code":
				text = "`" + text + "`"
			case "underlined":
				text = "<u>" + text + "</u>"
			case "strikethrough":
				text = "~~" + text + "~~"
			}
		}
		result.WriteString(text)
	}
	return result.String()
}

// serializeSlateNodes recursively processes nodes and its content into a plain-text format.
func serializeSlateNodes(nodes []SlateNode, nodeSeparator, leaveSeparator string) string {
	var result strings.Builder
	modifiedSlateNodeSeparator := nodeSeparator

	for _, node := range nodes {
		// Handle paragraph nodes by ensuring they end with punctuation.
		if node.Type == _nodeTypeParagraph {
			node.ensureEndsWithPunctuation()
		}

		// Recursively process child nodes.
		switch {
		case len(node.Nodes) > 0:
			switch node.Type {
			case _nodeTypeHeadingLarge, _nodeTypeCaption, _nodeTypeFigure:
				continue
			case _nodeTypeHeadingMedium:
				modifiedSlateNodeSeparator = _sentenceSeparator
			case _nodeTypeBulletedList, _nodeTypeNumberedList:
				node.Nodes[len(node.Nodes)-1].isLastInList = true
				temp := serializeSlateNodes(node.Nodes, modifiedSlateNodeSeparator, leaveSeparator)
				cleaned := cleanUpList(temp)
				result.WriteString(cleaned)
				continue
			case _nodeTypeInline:
				modifiedSlateNodeSeparator = _spaceSeparator
				result.WriteString(serializeSlateNodes(node.Nodes, modifiedSlateNodeSeparator, leaveSeparator) + leaveSeparator)
			case _nodeTypeLink:
				modifiedSlateNodeSeparator = _spaceSeparator
				result.WriteString(serializeSlateNodes(node.Nodes, modifiedSlateNodeSeparator, leaveSeparator))
			case _nodeTypeListItem:
				modifiedSlateNodeSeparator = _commaSeparator
				if node.isLastInList {
					modifiedSlateNodeSeparator = _sentenceSeparator
				}
				result.WriteString(serializeSlateNodes(node.Nodes, modifiedSlateNodeSeparator, leaveSeparator) + leaveSeparator)
			default:
				res := serializeSlateNodes(node.Nodes, _spaceSeparator, leaveSeparator)
				res = strings.TrimSpace(res)
				if modifiedSlateNodeSeparator == _commaSeparator && endsWithPunctuation(res) {
					res = res[:len(res)-1] + modifiedSlateNodeSeparator
				} else if modifiedSlateNodeSeparator == _newline {
					res += modifiedSlateNodeSeparator
				}
				result.WriteString(res)
			}
		case len(node.Leaves) > 0:
			result.WriteString(serializeSlateLeaves(node.Leaves, leaveSeparator))
		}
	}

	return result.String()
}

// serializeSlateLeaves joins leaf texts with the specified separator.
func serializeSlateLeaves(leaves []SlateLeaf, separator string) string {
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
func (node *SlateNode) ensureEndsWithPunctuation() {
	if len(node.Leaves) > 0 {
		lastSlateLeaf := node.Leaves[len(node.Leaves)-1]
		if text := lastSlateLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastSlateLeaf.Text += _sentenceSeparator
			}
		}
		node.Leaves[len(node.Leaves)-1] = lastSlateLeaf
	}

	if len(node.Nodes) > 0 && len(node.Nodes[len(node.Nodes)-1].Leaves) > 0 {
		lastSlateLeaf := node.Nodes[len(node.Nodes)-1].Leaves[len(node.Nodes[len(node.Nodes)-1].Leaves)-1]
		if text := lastSlateLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastSlateLeaf.Text += _sentenceSeparator
			}
		}
		node.Nodes[len(node.Nodes)-1].Leaves[len(node.Nodes[len(node.Nodes)-1].Leaves)-1] = lastSlateLeaf
	}
}

// cleans up the serialized list by removing excessive punctuation.
func cleanUpList(text string) string {
	cleaned := _multipleDotsRegex.ReplaceAllString(text, _sentenceSeparator)
	return _dotSpaceRegex.ReplaceAllString(cleaned, _dotSeparator)
}

func endsWithPunctuation(text string) bool {
	if len(text) == 0 {
		return false
	}

	lastChar := string(text[len(text)-1])
	return strings.Contains(_punctuationMarks, lastChar)
}
