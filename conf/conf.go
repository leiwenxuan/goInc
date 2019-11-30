package conf

const (
	CheckHead     = "/*--user=%s;--password=%s;--host=%s;--check=1;--port=%s;*/"
	ExecuteHead   = "/*--user=%s;--password=%s;--host=%s;--port=%s;--execute=1;--backup=1;ignore_warnings=1;*/"
	InceptionBody = `inception_magic_start;
			use %s;
			%s;
			inception_magic_commit;`
)
