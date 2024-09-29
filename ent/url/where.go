// Code generated by ent, DO NOT EDIT.

package enturl

import (
	"wsw/backend/domain/url"
	"wsw/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldID, id))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// RelativePath applies equality check predicate on the "relative_path" field. It's identical to RelativePathEQ.
func RelativePath(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldRelativePath, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Url {
	return predicate.Url(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Url {
	return predicate.Url(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Url {
	return predicate.Url(sql.FieldContainsFold(FieldURL, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v url.Status) predicate.Url {
	vc := v
	return predicate.Url(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v url.Status) predicate.Url {
	vc := v
	return predicate.Url(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...url.Status) predicate.Url {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...url.Status) predicate.Url {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(sql.FieldNotIn(FieldStatus, v...))
}

// RelativePathEQ applies the EQ predicate on the "relative_path" field.
func RelativePathEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldRelativePath, v))
}

// RelativePathNEQ applies the NEQ predicate on the "relative_path" field.
func RelativePathNEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldRelativePath, v))
}

// RelativePathIn applies the In predicate on the "relative_path" field.
func RelativePathIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldRelativePath, vs...))
}

// RelativePathNotIn applies the NotIn predicate on the "relative_path" field.
func RelativePathNotIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldRelativePath, vs...))
}

// RelativePathGT applies the GT predicate on the "relative_path" field.
func RelativePathGT(v string) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldRelativePath, v))
}

// RelativePathGTE applies the GTE predicate on the "relative_path" field.
func RelativePathGTE(v string) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldRelativePath, v))
}

// RelativePathLT applies the LT predicate on the "relative_path" field.
func RelativePathLT(v string) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldRelativePath, v))
}

// RelativePathLTE applies the LTE predicate on the "relative_path" field.
func RelativePathLTE(v string) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldRelativePath, v))
}

// RelativePathContains applies the Contains predicate on the "relative_path" field.
func RelativePathContains(v string) predicate.Url {
	return predicate.Url(sql.FieldContains(FieldRelativePath, v))
}

// RelativePathHasPrefix applies the HasPrefix predicate on the "relative_path" field.
func RelativePathHasPrefix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasPrefix(FieldRelativePath, v))
}

// RelativePathHasSuffix applies the HasSuffix predicate on the "relative_path" field.
func RelativePathHasSuffix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasSuffix(FieldRelativePath, v))
}

// RelativePathIsNil applies the IsNil predicate on the "relative_path" field.
func RelativePathIsNil() predicate.Url {
	return predicate.Url(sql.FieldIsNull(FieldRelativePath))
}

// RelativePathNotNil applies the NotNil predicate on the "relative_path" field.
func RelativePathNotNil() predicate.Url {
	return predicate.Url(sql.FieldNotNull(FieldRelativePath))
}

// RelativePathEqualFold applies the EqualFold predicate on the "relative_path" field.
func RelativePathEqualFold(v string) predicate.Url {
	return predicate.Url(sql.FieldEqualFold(FieldRelativePath, v))
}

// RelativePathContainsFold applies the ContainsFold predicate on the "relative_path" field.
func RelativePathContainsFold(v string) predicate.Url {
	return predicate.Url(sql.FieldContainsFold(FieldRelativePath, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Url) predicate.Url {
	return predicate.Url(sql.NotPredicates(p))
}
