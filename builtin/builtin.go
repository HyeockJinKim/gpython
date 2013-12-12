// Built-in functions
package builtin

import (
	"fmt"
	"github.com/ncw/gpython/compile"
	"github.com/ncw/gpython/py"
	"github.com/ncw/gpython/vm"
	"unicode/utf8"
)

const builtin_doc = `Built-in functions, exceptions, and other objects.

Noteworthy: None is the 'nil' object; Ellipsis represents '...' in slices.`

// Initialise the module
func init() {
	methods := []*py.Method{
		py.NewMethod("__build_class__", builtin___build_class__, 0, build_class_doc),
		py.NewMethod("__import__", builtin___import__, 0, import_doc),
		py.NewMethod("abs", builtin_abs, 0, abs_doc),
		// py.NewMethod("all", builtin_all, 0, all_doc),
		// py.NewMethod("any", builtin_any, 0, any_doc),
		// py.NewMethod("ascii", builtin_ascii, 0, ascii_doc),
		// py.NewMethod("bin", builtin_bin, 0, bin_doc),
		// py.NewMethod("callable", builtin_callable, 0, callable_doc),
		py.NewMethod("chr", builtin_chr, 0, chr_doc),
		py.NewMethod("compile", builtin_compile, 0, compile_doc),
		// py.NewMethod("delattr", builtin_delattr, 0, delattr_doc),
		// py.NewMethod("dir", builtin_dir, 0, dir_doc),
		// py.NewMethod("divmod", builtin_divmod, 0, divmod_doc),
		// py.NewMethod("eval", builtin_eval, 0, eval_doc),
		// py.NewMethod("exec", builtin_exec, 0, exec_doc),
		// py.NewMethod("format", builtin_format, 0, format_doc),
		py.NewMethod("getattr", builtin_getattr, 0, getattr_doc),
		// py.NewMethod("globals", builtin_globals, py.METH_NOARGS, globals_doc),
		py.NewMethod("hasattr", builtin_hasattr, 0, hasattr_doc),
		// py.NewMethod("hash", builtin_hash, 0, hash_doc),
		// py.NewMethod("hex", builtin_hex, 0, hex_doc),
		// py.NewMethod("id", builtin_id, 0, id_doc),
		// py.NewMethod("input", builtin_input, 0, input_doc),
		// py.NewMethod("isinstance", builtin_isinstance, 0, isinstance_doc),
		// py.NewMethod("issubclass", builtin_issubclass, 0, issubclass_doc),
		// py.NewMethod("iter", builtin_iter, 0, iter_doc),
		py.NewMethod("len", builtin_len, 0, len_doc),
		// py.NewMethod("locals", builtin_locals, py.METH_NOARGS, locals_doc),
		// py.NewMethod("max", builtin_max, 0, max_doc),
		// py.NewMethod("min", builtin_min, 0, min_doc),
		py.NewMethod("next", builtin_next, 0, next_doc),
		// py.NewMethod("oct", builtin_oct, 0, oct_doc),
		py.NewMethod("ord", builtin_ord, 0, ord_doc),
		py.NewMethod("pow", builtin_pow, 0, pow_doc),
		py.NewMethod("print", builtin_print, 0, print_doc),
		// py.NewMethod("repr", builtin_repr, 0, repr_doc),
		py.NewMethod("round", builtin_round, 0, round_doc),
		py.NewMethod("setattr", builtin_setattr, 0, setattr_doc),
		// py.NewMethod("sorted", builtin_sorted, 0, sorted_doc),
		// py.NewMethod("sum", builtin_sum, 0, sum_doc),
		// py.NewMethod("vars", builtin_vars, 0, vars_doc),
	}
	globals := py.StringDict{
		"None":     py.None,
		"Ellipsis": py.Ellipsis,
		"False":    py.False,
		"True":     py.True,
		"bool":     py.BoolType,
		// "memoryview":     py.MemoryViewType,
		// "bytearray":      py.ByteArrayType,
		"bytes":       py.BytesType,
		"classmethod": py.ClassMethodType,
		"complex":     py.ComplexType,
		"dict":        py.StringDictType, // FIXME
		// "enumerate":      py.EnumType,
		// "filter":         py.FilterType,
		"float":     py.FloatType,
		"frozenset": py.FrozenSetType,
		// "property":       py.PropertyType,
		"int":  py.IntType, // FIXME LongType?
		"list": py.ListType,
		// "map":            py.MapType,
		"object": py.ObjectType,
		"range":  py.RangeType,
		// "reversed":       py.ReversedType,
		"set": py.SetType,
		// "slice":          py.SliceType,
		"staticmethod": py.StaticMethodType,
		"str":          py.StringType,
		// "super":          py.SuperType,
		"tuple": py.TupleType,
		"type":  py.TypeType,
		// "zip":            py.ZipType,

		// Exceptions
		"ArithmeticError":           py.ArithmeticError,
		"AssertionError":            py.AssertionError,
		"AttributeError":            py.AttributeError,
		"BaseException":             py.BaseException,
		"BlockingIOError":           py.BlockingIOError,
		"BrokenPipeError":           py.BrokenPipeError,
		"BufferError":               py.BufferError,
		"BytesWarning":              py.BytesWarning,
		"ChildProcessError":         py.ChildProcessError,
		"ConnectionAbortedError":    py.ConnectionAbortedError,
		"ConnectionError":           py.ConnectionError,
		"ConnectionRefusedError":    py.ConnectionRefusedError,
		"ConnectionResetError":      py.ConnectionResetError,
		"DeprecationWarning":        py.DeprecationWarning,
		"EOFError":                  py.EOFError,
		"EnvironmentError":          py.OSError,
		"Exception":                 py.ExceptionType,
		"FileExistsError":           py.FileExistsError,
		"FileNotFoundError":         py.FileNotFoundError,
		"FloatingPointError":        py.FloatingPointError,
		"FutureWarning":             py.FutureWarning,
		"GeneratorExit":             py.GeneratorExit,
		"IOError":                   py.OSError,
		"ImportError":               py.ImportError,
		"ImportWarning":             py.ImportWarning,
		"IndentationError":          py.IndentationError,
		"IndexError":                py.IndexError,
		"InterruptedError":          py.InterruptedError,
		"IsADirectoryError":         py.IsADirectoryError,
		"KeyError":                  py.KeyError,
		"KeyboardInterrupt":         py.KeyboardInterrupt,
		"LookupError":               py.LookupError,
		"MemoryError":               py.MemoryError,
		"NameError":                 py.NameError,
		"NotADirectoryError":        py.NotADirectoryError,
		"NotImplemented":            py.NotImplemented,
		"NotImplementedError":       py.NotImplementedError,
		"OSError":                   py.OSError,
		"OverflowError":             py.OverflowError,
		"PendingDeprecationWarning": py.PendingDeprecationWarning,
		"PermissionError":           py.PermissionError,
		"ProcessLookupError":        py.ProcessLookupError,
		"ReferenceError":            py.ReferenceError,
		"ResourceWarning":           py.ResourceWarning,
		"RuntimeError":              py.RuntimeError,
		"RuntimeWarning":            py.RuntimeWarning,
		"StopIteration":             py.StopIteration,
		"SyntaxError":               py.SyntaxError,
		"SyntaxWarning":             py.SyntaxWarning,
		"SystemError":               py.SystemError,
		"SystemExit":                py.SystemExit,
		"TabError":                  py.TabError,
		"TimeoutError":              py.TimeoutError,
		"TypeError":                 py.TypeError,
		"UnboundLocalError":         py.UnboundLocalError,
		"UnicodeDecodeError":        py.UnicodeDecodeError,
		"UnicodeEncodeError":        py.UnicodeEncodeError,
		"UnicodeError":              py.UnicodeError,
		"UnicodeTranslateError":     py.UnicodeTranslateError,
		"UnicodeWarning":            py.UnicodeWarning,
		"UserWarning":               py.UserWarning,
		"ValueError":                py.ValueError,
		"Warning":                   py.Warning,
		"ZeroDivisionError":         py.ZeroDivisionError,
	}
	py.NewModule("builtins", builtin_doc, methods, globals)
}

const print_doc = `print(value, ..., sep=' ', end='\\n', file=sys.stdout, flush=False)

Prints the values to a stream, or to sys.stdout by default.
Optional keyword arguments:
file:  a file-like object (stream); defaults to the current sys.stdout.
sep:   string inserted between values, default a space.
end:   string appended after the last value, default a newline.
flush: whether to forcibly flush the stream.`

func builtin_print(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	fmt.Printf("print %v, %v, %v\n", self, args, kwargs)
	return py.None
}

const pow_doc = `pow(x, y[, z]) -> number

With two arguments, equivalent to x**y.  With three arguments,
equivalent to (x**y) % z, but may be more efficient (e.g. for ints).`

func builtin_pow(self py.Object, args py.Tuple) py.Object {
	var v, w, z py.Object
	z = py.None
	py.UnpackTuple(args, nil, "pow", 2, 3, &v, &w, &z)
	return py.Pow(v, w, z)
}

const abs_doc = `"abs(number) -> number

Return the absolute value of the argument.`

func builtin_abs(self, v py.Object) py.Object {
	return py.Abs(v)
}

const round_doc = `round(number[, ndigits]) -> number

Round a number to a given precision in decimal digits (default 0 digits).
This returns an int when called with one argument, otherwise the
same type as the number. ndigits may be negative.`

func builtin_round(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	var number, ndigits py.Object
	ndigits = py.Int(0)
	// var kwlist = []string{"number", "ndigits"}
	// FIXME py.ParseTupleAndKeywords(args, kwargs, "O|O:round", kwlist, &number, &ndigits)
	py.UnpackTuple(args, nil, "round", 1, 2, &number, &ndigits)

	numberRounder, ok := number.(py.I__round__)
	if !ok {
		panic(py.ExceptionNewf(py.TypeError, "type %s doesn't define __round__ method", number.Type().Name))
	}

	return numberRounder.M__round__(ndigits)
}

const build_class_doc = `__build_class__(func, name, *bases, metaclass=None, **kwds) -> class

Internal helper function used by the class statement.`

func builtin___build_class__(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	// fmt.Printf("__build_class__(self=%#v, args=%#v, kwargs=%#v\n", self, args, kwargs)
	var prep, cell, cls py.Object
	var mkw, ns py.StringDict
	var meta, winner *py.Type
	var isclass bool

	if len(args) < 2 {
		panic(py.ExceptionNewf(py.TypeError, "__build_class__: not enough arguments"))
	}

	// Better be callable
	fn, ok := args[0].(*py.Function)
	if !ok {
		panic(py.ExceptionNewf(py.TypeError, "__build__class__: func must be a function"))
	}

	name := args[1].(py.String)
	if !ok {
		panic(py.ExceptionNewf(py.TypeError, "__build_class__: name is not a string"))
	}
	bases := args[2:]

	if kwargs != nil {
		mkw = kwargs.Copy()      // Don't modify kwds passed in!
		meta := mkw["metaclass"] // _PyDict_GetItemId(mkw, &PyId_metaclass)
		if meta != nil {
			delete(mkw, "metaclass")
			// metaclass is explicitly given, check if it's indeed a class
			_, isclass = meta.(*py.Type)
		}
	}
	if meta == nil {
		// if there are no bases, use type:
		if len(bases) == 0 {
			meta = py.TypeType
		} else {
			// else get the type of the first base
			meta = bases[0].Type()
		}
		isclass = true // meta is really a class
	}

	if isclass {
		// meta is really a class, so check for a more derived
		// metaclass, or possible metaclass conflicts:
		winner = meta.CalculateMetaclass(bases)
		if winner != meta {
			meta = winner
		}
	}
	// else: meta is not a class, so we cannot do the metaclass
	// calculation, so we will use the explicitly given object as it is
	prep = meta.Type().Dict["___prepare__"] // FIXME should be using _PyObject_GetAttr
	if prep == nil {
		ns = py.NewStringDict()
	} else {
		ns = py.Call(prep, py.Tuple{name, bases}, mkw).(py.StringDict)
	}
	// fmt.Printf("Calling %v with %p and %p\n", fn.Name, fn.Globals, ns)
	// fmt.Printf("Code = %#v\n", fn.Code)
	locals := fn.LocalsForCall(py.Tuple{ns})
	cell, err := vm.Run(fn.Globals, locals, fn.Code, fn.Closure)

	// fmt.Printf("result = %#v err = %s\n", cell, err)
	// fmt.Printf("locals = %#v\n", locals)
	// fmt.Printf("ns = %#v\n", ns)
	if err != nil {
		// propagate the error
		panic(err)
	}
	if cell != nil {
		// fmt.Printf("Calling %v\n", meta)
		cls = py.Call(meta, py.Tuple{name, bases, ns}, mkw)
		if c, ok := cell.(*py.Cell); ok {
			c.Set(cls)
		}
	}
	// fmt.Printf("Globals = %v, Locals = %v\n", fn.Globals, ns)
	return cls
}

const next_doc = `next(iterator[, default])

Return the next item from the iterator. If default is given and the iterator
is exhausted, it is returned instead of raising StopIteration.`

func builtin_next(self py.Object, args py.Tuple) (res py.Object) {
	var it, def py.Object

	py.UnpackTuple(args, nil, "next", 1, 2, &it, &def)

	if def != nil {
		defer func() {
			if r := recover(); r != nil {
				if py.IsException(py.StopIteration, r) {
					// Return defult on StopIteration
					res = def
				} else {
					// Re-raise
					panic(r)
				}
			}
		}()
	}

	return py.Next(it)
}

const import_doc = `__import__(name, globals=None, locals=None, fromlist=(), level=0) -> module

Import a module. Because this function is meant for use by the Python
interpreter and not for general use it is better to use
importlib.import_module() to programmatically import a module.

The globals argument is only used to determine the context;
they are not modified.  The locals argument is unused.  The fromlist
should be a list of names to emulate ''from name import ...'', or an
empty list to emulate ''import name''.
When importing a module from a package, note that __import__('A.B', ...)
returns package A when fromlist is empty, but its submodule B when
fromlist is not empty.  Level is used to determine whether to perform 
absolute or relative imports. 0 is absolute while a positive number
is the number of parent directories to search relative to the current module.`

func builtin___import__(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	kwlist := []string{"name", "globals", "locals", "fromlist", "level"}
	var name py.Object
	var globals py.Object = py.NewStringDict()
	var locals py.Object = py.NewStringDict()
	var fromlist py.Object = py.Tuple{}
	var level py.Object = py.Int(0)

	py.ParseTupleAndKeywords(args, kwargs, "U|OOOi:__import__", kwlist, &name, &globals, &locals, &fromlist, &level)
	if fromlist == py.None {
		fromlist = py.Tuple{}
	}
	return py.ImportModuleLevelObject(string(name.(py.String)), globals.(py.StringDict), locals.(py.StringDict), fromlist.(py.Tuple), int(level.(py.Int)))
}

const ord_doc = `ord(c) -> integer

Return the integer ordinal of a one-character string.`

func builtin_ord(self, obj py.Object) py.Object {
	var size int
	switch x := obj.(type) {
	case py.Bytes:
		size = len(x)
		if len(x) == 1 {
			return py.Int(x[0])
		}
	case py.String:
		var rune rune
		rune, size = utf8.DecodeRuneInString(string(x))
		if len(x) == size && rune != utf8.RuneError {
			return py.Int(rune)
		}
	//case py.ByteArray:
	// XXX Hopefully this is temporary
	// FIXME implement
	// size = PyByteArray_GET_SIZE(obj)
	// if size == 1 {
	// 	ord = (long)((char) * PyByteArray_AS_STRING(obj))
	// 	return PyLong_FromLong(ord)
	// }
	default:
		panic(py.ExceptionNewf(py.TypeError, "ord() expected string of length 1, but %s found", obj.Type().Name))
	}

	panic(py.ExceptionNewf(py.TypeError, "ord() expected a character, but string of length %zd found", size))
}

const getattr_doc = `getattr(object, name[, default]) -> value

Get a named attribute from an object; getattr(x, 'y') is equivalent to x.y.
When a default argument is given, it is returned when the attribute doesn't
exist; without it, an exception is raised in that case.`

func builtin_getattr(self py.Object, args py.Tuple) py.Object {
	var v, result, dflt py.Object
	var name py.Object

	py.UnpackTuple(args, nil, "getattr", 2, 3, &v, &name, &dflt)

	nameStr, ok := name.(py.String)
	if !ok {
		panic(py.ExceptionNewf(py.TypeError, "getattr(): attribute name must be string"))
	}

	result, err := py.GetAttrErr(v, string(nameStr))
	if err != nil {
		if dflt == nil {
			panic(err)
		}
		result = dflt
	}
	return result
}

const hasattr_doc = `hasattr(object, name) -> bool

Return whether the object has an attribute with the given name.
(This is done by calling getattr(object, name) and catching AttributeError.)`

func builtin_hasattr(self py.Object, args py.Tuple) py.Object {
	var v py.Object
	var name py.Object
	py.UnpackTuple(args, nil, "hasattr", 2, 2, &v, &name)

	nameStr, ok := name.(py.String)
	if !ok {
		panic(py.ExceptionNewf(py.TypeError, "hasattr(): attribute name must be string"))
	}

	_, err := py.GetAttrErr(v, string(nameStr))
	return py.NewBool(err == nil)
}

const setattr_doc = `setattr(object, name, value)

Set a named attribute on an object; setattr(x, 'y', v) is equivalent to
"x.y = v".`

func builtin_setattr(self py.Object, args py.Tuple) py.Object {
	var v py.Object
	var name py.Object
	var value py.Object

	py.UnpackTuple(args, nil, "setattr", 3, 3, &v, &name, &value)

	return py.SetAttr(v, name, value)
}

// Reads the source as a string
func source_as_string(cmd py.Object, funcname, what string /*, PyCompilerFlags *cf */) string {
	// FIXME only understands strings, not bytes etc at the moment
	if str, ok := cmd.(py.String); ok {
		// FIXME cf->cf_flags |= PyCF_IGNORE_COOKIE;
		return string(str)
	}
	// } else if (!PyObject_CheckReadBuffer(cmd)) {
	panic(py.ExceptionNewf(py.TypeError, "%s() arg 1 must be a %s object", funcname, what))
	// } else if (PyObject_AsReadBuffer(cmd, (const void **)&str, &size) < 0) {
	// 	return nil;
}

const compile_doc = `compile(source, filename, mode[, flags[, dont_inherit]]) -> code object

Compile the source string (a Python module, statement or expression)
into a code object that can be executed by exec() or eval().
The filename will be used for run-time error messages.
The mode must be 'exec' to compile a module, 'single' to compile a
single (interactive) statement, or 'eval' to compile an expression.
The flags argument, if present, controls which future statements influence
the compilation of the code.
The dont_inherit argument, if non-zero, stops the compilation inheriting
the effects of any future statements in effect in the code calling
compile; if absent or zero these statements do influence the compilation,
in addition to any features explicitly specified.`

func builtin_compile(self py.Object, args py.Tuple, kwargs py.StringDict) py.Object {
	// FIXME lots of unsupported stuff here!
	var filename py.Object
	var startstr py.Object
	// var mode = -1
	var dont_inherit py.Object = py.Int(0)
	var supplied_flags py.Object = py.Int(0)
	var optimizeInt py.Object = py.Int(-1)
	//is_ast := false
	// var cf PyCompilerFlags
	var cmd py.Object
	kwlist := []string{"source", "filename", "mode", "flags", "dont_inherit", "optimize"}
	// start := []int{Py_file_input, Py_eval_input, Py_single_input}
	var result py.Object

	py.ParseTupleAndKeywords(args, kwargs, "Oss|iii:compile", kwlist,
		&cmd,
		&filename,
		&startstr,
		&supplied_flags,
		&dont_inherit,
		&optimizeInt)

	// cf.cf_flags = supplied_flags | PyCF_SOURCE_IS_UTF8

	// if supplied_flags&^(PyCF_MASK|PyCF_MASK_OBSOLETE|PyCF_DONT_IMPLY_DEDENT|PyCF_ONLY_AST) != 0 {
	// 	panic(py.ExceptionNewf(py.ValueError, "compile(): unrecognised flags"))
	// }
	// XXX Warn if (supplied_flags & PyCF_MASK_OBSOLETE) != 0?

	optimize := int(optimizeInt.(py.Int))
	if optimize < -1 || optimize > 2 {
		panic(py.ExceptionNewf(py.ValueError, "compile(): invalid optimize value"))
	}

	if dont_inherit.(py.Int) != 0 {
		// PyEval_MergeCompilerFlags(&cf)
	}

	// switch string(startstr.(py.String)) {
	// case "exec":
	// 	mode = 0
	// case "eval":
	// 	mode = 1
	// case "single":
	// 	mode = 2
	// default:
	// 	panic(py.ExceptionNewf(py.ValueError, "compile() arg 3 must be 'exec', 'eval' or 'single'"))
	// }

	// is_ast = PyAST_Check(cmd)
	// if is_ast {
	// 	if supplied_flags & PyCF_ONLY_AST {
	// 		result = cmd
	// 	} else {

	// 		arena := PyArena_New()
	// 		mod := PyAST_obj2mod(cmd, arena, mode)
	// 		PyAST_Validate(mod)
	// 		result = PyAST_CompileObject(mod, filename, &cf, optimize, arena)
	// 		PyArena_Free(arena)
	// 	}
	// } else {
	str := source_as_string(cmd, "compile", "string, bytes or AST" /*, &cf*/)
	// result = py.CompileStringExFlags(str, filename, start[mode], &cf, optimize)
	result = compile.Compile(str, string(filename.(py.String)), string(startstr.(py.String)), int(supplied_flags.(py.Int)), dont_inherit.(py.Int) != 0)
	// }

	return result
}

const len_doc = `len(object) -> integer

Return the number of items of a sequence or mapping.`

func builtin_len(self, v py.Object) py.Object {
	return py.Len(v)
}

const chr_doc = `chr(i) -> Unicode character

Return a Unicode string of one character with ordinal i; 0 <= i <= 0x10ffff.`

func builtin_chr(self py.Object, args py.Tuple) py.Object {
	var xObj py.Object

	py.ParseTuple(args, "i:chr", &xObj)

	x := xObj.(py.Int)
	if x < 0 || x >= 0x110000 {
		panic(py.ExceptionNewf(py.ValueError, "chr() arg not in range(0x110000)"))
	}
	buf := make([]byte, 8)
	n := utf8.EncodeRune(buf, rune(x))
	return py.String(buf[:n])
}
