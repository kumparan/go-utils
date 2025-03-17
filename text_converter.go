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
)

// Leaf represents a text element with optional formatting.
type Leaf struct {
	Object string `json:"object"`
	Text   string `json:"text"`
	Marks  []any  `json:"marks"`
}

// Node represents a hierarchical document structure.
type Node struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Nodes  []Node `json:"nodes,omitempty"`
	Leaves []Leaf `json:"leaves,omitempty"`

	isLastListElement bool
}

// Document represents the root structure of a Slate document.
type Document struct {
	Nodes []Node `json:"nodes"`
}

// WrappedDocument is used to deserialize Slate JSON documents.
type WrappedDocument struct {
	Document Document `json:"document"`
}

// SerializeSlateToText processes and formats a Slate document JSON string.
func SerializeSlateToText(documentJSON string) (string, error) {
	var wrappedDocument WrappedDocument
	if err := json.Unmarshal([]byte(documentJSON), &wrappedDocument); err != nil {
		log.WithField("documentJSON", documentJSON).Error(err)
		return "", err
	}

	text := serializeNodes(wrappedDocument.Document.Nodes, _newline, _spaceSeparator)
	text = regexp.MustCompile(`\n+`).ReplaceAllString(text, _newline)

	return strings.TrimSpace(text), nil
}

// serializeNodes recursively processes nodes and their content into a plain-text format.
func serializeNodes(nodes []Node, nodeSeparator, leaveSeparator string) string {
	var result strings.Builder
	modifiedNodeSeparator := nodeSeparator

	for _, node := range nodes {
		// Handle paragraph nodes by ensuring they end with punctuation.
		if node.Type == "paragraph" {
			ensureEndsWithPunctuation(&node)
		}

		// Recursively process child nodes.
		if len(node.Nodes) > 0 && node.Type != "heading-large" && node.Type != "caption" && node.Type != "figure" {
			// modifiedNodeSeparator = getModifiedNodeSeparator(node.Type, nodeSeparator)
			switch node.Type {
			case "heading-medium":
				return _sentenceSeparator
			case "bulleted-list", "numbered-list":
				node.Nodes[len(node.Nodes)-1].isLastListElement = true
				temp := serializeNodes(node.Nodes, modifiedNodeSeparator, leaveSeparator)
				cleaned := cleanUpList(temp)
				result.WriteString(cleaned)
				continue
			case "inline":
				modifiedNodeSeparator = _spaceSeparator
				result.WriteString(serializeNodes(node.Nodes, modifiedNodeSeparator, leaveSeparator) + leaveSeparator)
			case "list-item":
				modifiedNodeSeparator = _commaSeparator
				if node.isLastListElement {
					modifiedNodeSeparator = ". "
				}
				result.WriteString(serializeNodes(node.Nodes, modifiedNodeSeparator, leaveSeparator) + leaveSeparator)
			default:
				res := serializeNodes(node.Nodes, " ", leaveSeparator)
				res = strings.TrimSpace(res)
				if modifiedNodeSeparator == _commaSeparator && endsWithPunctuation(res) {
					res = res[:len(res)-1] + modifiedNodeSeparator
				} else if modifiedNodeSeparator == _newline {
					res += modifiedNodeSeparator
				}
				result.WriteString(res)
			}
		} else if len(node.Leaves) > 0 {
			result.WriteString(serializeLeaves(node.Leaves, leaveSeparator))
		}
	}

	return result.String()
}

// serializeLeaves joins leaf texts with the specified separator.
func serializeLeaves(leaves []Leaf, separator string) string {
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
func ensureEndsWithPunctuation(node *Node) {
	if len(node.Leaves) > 0 {
		lastLeaf := node.Leaves[len(node.Leaves)-1]
		if text := lastLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastLeaf.Text += _sentenceSeparator
			}
		}
		node.Leaves[len(node.Leaves)-1] = lastLeaf
	}

	if len(node.Nodes) > 0 && len(node.Nodes[len(node.Nodes)-1].Leaves) > 0 {
		lastLeaf := node.Nodes[len(node.Nodes)-1].Leaves[len(node.Nodes[len(node.Nodes)-1].Leaves)-1]
		if text := lastLeaf.Text; text != "" {
			if !endsWithPunctuation(text) {
				lastLeaf.Text += _sentenceSeparator
			}
		}
		node.Nodes[len(node.Nodes)-1].Leaves[len(node.Nodes[len(node.Nodes)-1].Leaves)-1] = lastLeaf
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
