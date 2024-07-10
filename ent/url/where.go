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

// APIURLID applies equality check predicate on the "api_url_id" field. It's identical to APIURLIDEQ.
func APIURLID(v int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldAPIURLID, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldImage, v))
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

// APIURLIDEQ applies the EQ predicate on the "api_url_id" field.
func APIURLIDEQ(v int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldAPIURLID, v))
}

// APIURLIDNEQ applies the NEQ predicate on the "api_url_id" field.
func APIURLIDNEQ(v int) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldAPIURLID, v))
}

// APIURLIDIn applies the In predicate on the "api_url_id" field.
func APIURLIDIn(vs ...int) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldAPIURLID, vs...))
}

// APIURLIDNotIn applies the NotIn predicate on the "api_url_id" field.
func APIURLIDNotIn(vs ...int) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldAPIURLID, vs...))
}

// APIURLIDGT applies the GT predicate on the "api_url_id" field.
func APIURLIDGT(v int) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldAPIURLID, v))
}

// APIURLIDGTE applies the GTE predicate on the "api_url_id" field.
func APIURLIDGTE(v int) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldAPIURLID, v))
}

// APIURLIDLT applies the LT predicate on the "api_url_id" field.
func APIURLIDLT(v int) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldAPIURLID, v))
}

// APIURLIDLTE applies the LTE predicate on the "api_url_id" field.
func APIURLIDLTE(v int) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldAPIURLID, v))
}

// APIURLIDIsNil applies the IsNil predicate on the "api_url_id" field.
func APIURLIDIsNil() predicate.Url {
	return predicate.Url(sql.FieldIsNull(FieldAPIURLID))
}

// APIURLIDNotNil applies the NotNil predicate on the "api_url_id" field.
func APIURLIDNotNil() predicate.Url {
	return predicate.Url(sql.FieldNotNull(FieldAPIURLID))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Url {
	return predicate.Url(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Url {
	return predicate.Url(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Url {
	return predicate.Url(sql.FieldContainsFold(FieldImage, v))
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
