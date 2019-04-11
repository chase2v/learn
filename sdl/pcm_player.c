#include <SDL.h>
#define BLOCK_SIZE 4096000

static Unit8 *audio_buf = NULL;

int main(int argc, char* argv[])
{
	int ret = -1;
	char *path = "./1.pcm";
	FILE *audio_fd = NULL;

	if (SDL_Init(SDL_INIT_AUDIO)) {
		SDL_Log("Failed to initial!");
		return ret;
	}

	audio_fd = fopen(path, "r");
	if (!audio_fd) {
		SDL_Log("Failed to open pcm file!");
		goto __FAIL;
	}

	audio_buf = (Unit8*)malloc(BLOCK_SIZE);

__FAIL:
	SDL_Quit();

	return 0;
}
