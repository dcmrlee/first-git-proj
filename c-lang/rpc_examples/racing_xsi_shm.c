#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <fcntl.h>
#include <string.h>
#include <sys/file.h>
#include <wait.h>
#include <sys/mman.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/types.h>

#define COUNT 100
#define PATHNAME "/etc/passwd"

int do_child(int proj_id)
{
	int interval;
	int *shm_p, shm_id;
	key_t shm_key;

	/* 使用ftok产生shmkey */
	if ((shm_key = ftok(PATHNAME, proj_id)) == -1) {
		perror("ftok()");
		exit(1);
	}

	/* 在子进程中使用shmget取得已经在父进程中创建好的共享内存 */
	shm_id = shmget(shm_key, sizeof(int), 0);
	if (shm_id < 0) {
		perror("shmget()");
		exit(1);
	}


	/* 使用shmat将相关恭喜内存段映射到本进程的内存地址 */
	shm_p = (int *)shmat(shm_id, NULL, 0);
	if ((void*)shm_p == (void*)-1) {
		perror("shmat()");
		exit(1);
	}

	/* critical section */
	interval = *shm_p;
	interval++;
	usleep(1);
	*shm_p = interval;
	/* critical section */


	/* 使用shmdt接触本进程内对共享内存的地址映射，本操作不会删除共享内存 */
	if (shmdt(shm_p) < 0) {
		perror("shmdt");
		exit(1);
	}
	exit(0);
}


int main()
{
	pid_t pid;
	int count;
	int *shm_p;
	int shm_id, proj_id;
	key_t shm_key;


	/* 使用约定好的文件路径和proj_id产生shm_key */
	if ((shm_key = ftok(PATHNAME, proj_id)) == -1) {
		perror("shm_key");
		exit(1);
	}

	/*
	* 使用shm_key创建一个共享内存
	* 如该系统中已经存在此共享内存则报错退出
	* 创建出来的共享内存权限为0600
	*/
	shm_id = shmget(shm_key, sizeof(int), IPC_CREAT|IPC_EXCL|0600);
	if (shm_id < 0) {
		perror("shm_get()");
		exit(1);
	}

	
	/* 将创建好的共享内存映射到父进程的地址以便访问 */
	shm_p = (int *)shmat(shm_id, NULL, 0);
	if ((void*)shm_p == (void*)-1) {
		perror("shmat()");
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
			do_child(proj_id);
		}
	}

	for (count=0; count<COUNT; count++) {
		wait(NULL);
	}

	printf("shm_p: %d\n", *shm_p);

	/* 解除共享内存地址映射 */
	if (shmdt(shm_p) < 0) {
		perror("shmdt");
		exit(1);
	}

	/* 删除共享内存 */
	if (shmctl(shm_id, IPC_RMID, NULL) < 0) {
		perror("shmctl()");
		exit(1);
	}

	exit(0);
}
