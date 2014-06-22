
#include "runtime.h"
#include "arch_GOARCH.h"
#include "defs_GOOS_GOARCH.h"
#include "os_GOOS.h"
#include "malloc.h"

void*
runtime·SysAlloc(uintptr n, uint64 *stat)
{
	runtime·throw("runtime: SysAlloc not implemented on freertos");
	return nil;
}

void
runtime·SysUnused(void *v, uintptr n)
{
	USED(v);
	USED(n);
}

void
runtime·SysUsed(void *v, uintptr n)
{
	USED(v);
	USED(n);
}

void
runtime·SysFree(void *v, uintptr n, uint64 *stat)
{
	runtime·throw("runtime: SysFree not implemented on freertos");
}

void
runtime·SysFault(void *v, uintptr n)
{
	runtime·throw("runtime: SysFault not implemented on freertos");
}

void*
runtime·SysReserve(void *v, uintptr n, bool *reserved)
{
	runtime·throw("runtime: SysReserve not implemented on freertos");
	return nil;
}

void
runtime·SysMap(void *v, uintptr n, bool reserved, uint64 *stat)
{
	runtime·throw("runtime: SysMap not implemented on freertos");
}
