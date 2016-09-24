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

int do_child(int *count)
{
	int interval;

	/* critical section */
	interval = *count;
	interval++;
	usleep(1);
	*count = interval;
	/* critical section */

	exit(0);
}


int main()
{
	pid_t pid;
	int count;
	int *shm_p;
	
	shm_p = (int *)mmap(NULL, sizeof(int), PROT_WRITE|PROT_READ,
						MAP_SHARED|MAP_ANONYMOUS, -1, 0);
	if (MAP_FAILED == shm_p) {
		perror("mmap()");
		exit(1);
	}

	*shm_p = 0;
	
	for (count=0; count<COUNT; count++) {
		pid = fork();
		if (pid < 0) {
			perror("fork()");
			exit(1);
		}

		if (pid == 0) {
			do_child(shm_p);
		}
	}

	for (count=0; count<COUNT; count++) {
		wait(NULL);
	}

	printf("shm_p: %d\n", *shm_p);
	munmap(shm_p, sizeof(int));
	exit(0);
}
