const gulp = require('gulp');
const execFile = require('child_process').execFile;

const goFiles = './*.go';

gulp.task('build', () => {
	execFile(`${__dirname}/build.sh`, (error, stdout, stderr) => {
    if (error) {
      console.error(error);
    } else {
      if (stdout.length > 0) console.log(stdout);
			if (stderr) console.error(stderr);
    }
  });
});

gulp.task('default', [ 'build' ], () => {
	gulp.watch(goFiles, [ 'build' ]);
});
