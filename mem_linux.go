// +build linux

package mem

/*
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>

static long get_mem_available()
{
	FILE* fp = fopen( "/proc/meminfo", "r" );
	if ( fp != NULL )
	{
		size_t bufsize = 1024 * sizeof(char);
		char* buf      = (char*)malloc( bufsize );
		long value     = -1L;
		while ( getline( &buf, &bufsize, fp ) >= 0 )
		{
			//if ( strncmp( buf, "MemTotal", 8 ) != 0 )
			if ( strncmp( buf, "MemFree", 7 ) != 0 )
				continue;
			sscanf( buf, "%*s%ld", &value );
			break;
		}
		fclose( fp );
		free( (void*)buf );
		if ( value != -1L )
			return (size_t)(value * 1024L );
	}
	return 0;
}
	
static long get_mem_total()
{
	long pages = sysconf(_SC_PHYS_PAGES);
    long page_size = sysconf(_SC_PAGE_SIZE);
    return pages * page_size;
}
*/
import "C"

func GetAvailable() uint64 {
	avail := uint64(C.get_mem_available())
	return avail
}

func GetTotal() uint64 {
	total := uint64(C.get_mem_total())
	return total
}
