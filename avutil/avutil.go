/*
	The libavutil library is a utility library to aid portable multimedia programming.
	It contains safe portable string functions, random number generators, data structures,
	additional mathematics functions, cryptography and multimedia related functionality.
	Some generic features and utilities provided by the libavutil library
*/
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	Options       C.struct_AVOptions
	Dictionary    C.struct_AVDictionary
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	File          C.FILE
)

//unsigned 	avutil_version (void)
//Return the LIBAvUTIL_VERSION_INT constant.
func AvutilVersion() uint {
	return uint(C.avutil_version())
}

//const char * 	avutil_configuration (void)
//Return the libavutil build-time configuration.
func Avutil_configuration() string {
	return C.GoString(C.avutil_configuration())
}

//const char * 	avutil_license (void)
//Return the libavutil license.
func Avutil_license() string {
	return C.GoString(C.avutil_license())
}

//const char * 	av_get_media_type_string (enum MediaType media_type)
//Return a string describing the media_type enum, NULL if media_type is unknown.
func Av_get_media_type_string(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

//char av_get_picture_type_char (enum AvPictureType pict_type)
//Return a single letter to describe the given picture type pict_type.
func Av_get_picture_type_char(pt AvPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

//static void * av_x_if_null (const void *p, const void *x)
//Return x default pointer in case p is NULL.
func Av_x_if_null(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

//unsigned 	av_int_list_length_for_size (unsigned elsize, const void *list, uint64_t term) av_pure
//Compute the length of an integer list.
func Av_int_list_length_for_size(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

//FILE * 	av_fopen_utf8 (const char *path, const char *mode)
//Open a file using a UTF-8 filename.
func Av_fopen_utf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

//Rational 	av_get_time_base_q (void)
//Return the fractional representation of the internal time base.
func Av_get_time_base_q() Rational {
	return (Rational)(C.av_get_time_base_q())
}
