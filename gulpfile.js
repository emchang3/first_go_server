const gulp = require('gulp');
const execFile = require('child_process').execFile;

const goFiles = './*.go';

const restartServer = () => {
	execFile(`${__dirname}/restart_server.sh`, (error, stdout, stderr) => {
		if (error) {
			console.error(error);
		}
		else {
			if (stdout.length > 0) console.log(stdout);
			if (stderr) console.log(stderr);
		}
	});
}

const build = () => {
	execFile(`${__dirname}/build.sh`, (error, stdout, stderr) => {
    if (error) {
      console.error(error);
    } else {
      if (stdout.length > 0) {
				const buildStatus = stdout.split('')[0];
				if (buildStatus === '0') {
					restartServer();
				}
				else {
					if (stderr) console.log(stderr);
				}
			}
    }
  });
}

gulp.task('build', build);

gulp.task('default', [ 'build' ], () => {
	gulp.watch(goFiles, [ 'build' ]);
});
