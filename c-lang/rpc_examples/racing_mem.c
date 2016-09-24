#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <fcntl.h>
#include <string.h>
#include <sys/file.h>
#include <wait.h>
#include <sys/mman.h>

#define COUNT 100
#define MEMSIZE 1024*1024*1023*2


int main()
{
	pid_t pid;
	int count;
	int *shm_p;
	
	shm_p = (int *)mmap(NULL, MEMSIZE, PROT_WRITE|PROT_READ,
						MAP_SHARED|MAP_ANONYMOUS, -1, 0);
	if (MAP_FAILED == shm_p) {
		perror("mmap()");
		exit(1);
	}

	bzero(shm_p, MEMSIZE);

	sleep(3000);

	munmap(shm_p, MEMSIZE);
	exit(0);
}
