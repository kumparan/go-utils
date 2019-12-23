package utils

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
)

// Int64ToGraphQLID :nodoc:
func Int64ToGraphQLID(i int64) graphql.ID {
	return graphql.ID(Int64ToString(i))
}

// GraphQLIDToString :nodoc:
func GraphQLIDToString(id graphql.ID) string {
	return string(id)
}

// GraphQLIDPointerToString :nodoc:
func GraphQLIDPointerToString(id *graphql.ID) string {
	if id == nil {
		return ""
	}

	return string(*id)
}

// GraphQLIDToInt64 :nodoc:
func GraphQLIDToInt64(id graphql.ID) int64 {
	return StringToInt64(string(id))
}

// GraphQLIDPointerToInt64 :nodoc:
func GraphQLIDPointerToInt64(id *graphql.ID) int64 {
	if id == nil {
		return int64(0)
	}

	return StringToInt64(string(*id))
}

// GraphQLIDToInt32 :nodoc:
func GraphQLIDToInt32(id graphql.ID) int32 {
	newID, err := strconv.Atoi(string(id))
	if err != nil {
		return int32(0)
	}
	return int32(newID)
}

// GraphQLIDPointerToInt32 :nodoc:
func GraphQLIDPointerToInt32(id *graphql.ID) int32 {
	if id == nil {
		return int32(0)
	}

	newID, err := strconv.Atoi(string(*id))
	if err != nil {
		return int32(0)
	}
	return int32(newID)
}
