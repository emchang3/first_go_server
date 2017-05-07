require('dotenv').load();

const execFile = require('child_process').execFile;

const gulp = require('gulp');
const uglify = require('gulp-uglify');
const pump = require('pump');

const src = `${__dirname}/src`;
const intermediate = `${__dirname}/intermediate`;
const dest = `${__dirname}/public/scripts`;
const goFiles = './*.go';

const platform = process.platform === 'darwin' ? 'osx' : 'ubuntu'

const start = () => execFile(`${__dirname}/shell_ops_${platform}.sh`, [ 3 ]);

const started = () => execFile(
	`${__dirname}/shell_ops_${platform}.sh`,
	[ 4 ],
	(error, stdout, stderr) => {
		if (stdout.length > 0) {
			const going = stdout.split('\n')[0];
			if (going !== 'NOGO') {
				console.log(`--- STARTED: ${going} listening on :${process.env.PORT} ---\n`);
			}
		}
		if (stderr) console.log('sd_stderr:', stderr);
	}
);


const build = () => execFile(
	`${__dirname}/shell_ops_${platform}.sh`,
	[ 2 ],
	(error, stdout, stderr) => {
		if (error) {
			console.error('b_error', error);
		} else {
			if (stdout.length > 0) {
				const buildStatus = stdout.split('\n')[0];
				if (buildStatus === '0') {
					console.log('--- BUILD SUCCESSFUL ---');
					start();
					started();
				}
			}
			if (stderr) console.log('b_stderr', stderr);
		}
	}
);

gulp.task('build', [ 'destroy' ], build);

const destroy = () => {
	execFile(
		`${__dirname}/shell_ops_${platform}.sh`,
		[ 1 ],
		(error, stdout, stderr) => {
			if (stdout.length > 0) {
				const gopid = stdout.split('\n')[0];
				if (gopid !== 'NOGO') {
					console.log(`--- KILLED: ${gopid} ---`);
				} else {
					console.log('--- NO GO INSTANCE ---');
				}
			}
			if (stderr) console.log('d_stderr:', stderr);
		}
	);
}

gulp.task('destroy', destroy);

const cleanup = () => execFile(`${__dirname}/removeIntermediate.sh`, [ intermediate ]);

const minify = () => pump(
  [
    gulp.src(`${intermediate}/*.js`),
    uglify(),
    gulp.dest(dest)
  ],
  () => cleanup()
);

gulp.task(
  'compile',
  () => execFile(
    `${__dirname}/compileTS.sh`,
    [ `${src}/*.ts`, intermediate ],
    (error, stdout, stderr) => {
      // if (error) console.log(error);

      // NOTE: `./intermediate` will still exist if compile failed. Good visual cue in editor.
      stdout.length > 0 ? console.log(stdout) : minify()
    }
  )
);

gulp.task('default', [ 'build', 'compile' ], () => {
	gulp.watch(goFiles, [ 'build' ]);
	gulp.watch([ `${src}/*.ts` ], [ 'compile' ]);
});
