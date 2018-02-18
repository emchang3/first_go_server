namespace Eleven {
    {/* <div id="logo-box">
        <div id="j" class="j-out"></div>
        <div id="e-dash" class="e-hidden"></div>
        <div id="n" class="n-itial"></div>
        <div id="n-slash" class="n-itial-slash"></div>
        <div id="s1" class="s1-init"></div>
        <div id="s2" class="s2-init"></div>
        <div id="s3" class="s3-init"></div>
        <div id="q1" class="q1-init"></div>
        <div id="q2" class="q2-init"></div>
    </div> */}

    let state = 'jnsq';

    const eDash = document.getElementById('e-dash');
    const j = document.getElementById('j');
    const n = document.getElementById('n');
    const nSlash = document.getElementById('n-slash');
    const s1 = document.getElementById('s1');
    const s2 = document.getElementById('s2');
    const s3 = document.getElementById('s3');
    const q1 = document.getElementById('q1');
    const q2 = document.getElementById('q2');

    const change = () => {
        if (state == 'jnsq') {
            j.className = 'j-in';
            nSlash.className = 'n-d-slash';
            s1.className = 's1-inter';
            s2.className = 's2-inter';
            s3.className = 's3-inter';
            q1.className = 'q1-inter';
            q2.className = 'q2-inter';
            setTimeout(() => {
                j.className = 'e';
                eDash.className = 'e-shown';
                n.className = 'n-d';
                s1.className = 's1-end';
                s2.className = 's2-end';
                s3.className = 's3-end';
                q1.className = 'q1-end';
                q2.className = 'q2-end';
            }, 200);
            state = 'ezra';
        } else {
            eDash.className = 'e-hidden';
            j.className = 'j-in';
            nSlash.className = 'n-itial-slash';
            s1.className = 's1-inter';
            s2.className = 's2-inter';
            s3.className = 's3-inter';
            q1.className = 'q1-inter';
            q2.className = 'q2-inter';
            setTimeout(() => {
                j.className = 'j-out';
                n.className = 'n-itial';
                s1.className = 's1-init';
                s2.className = 's2-init';
                s3.className = 's3-init';
                q1.className = 'q1-init';
                q2.className = 'q2-init';
            }, 200);
            state = 'jnsq';
        }
    }

    setInterval(change, 2000);

    // j.addEventListener('click', change);
}