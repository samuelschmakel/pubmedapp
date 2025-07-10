const form = document.getElementById('myForm');

function init() {
    form.addEventListener('submit', handleClick);
}

// TODO: fix paper url link to go to link on click
async function handleClick(e) {
    console.log(`This event: ${e} happened.`);
    e.preventDefault(); // Prevent page reload

    const query = document.getElementById("query");

    // TODO: Make this less hideous than the alert box
    if (!query.value.trim()) {
        alert('The query field must not be empty.')
        return;
    }
    
    const formData = new FormData(e.target);
    for (const pair of formData.entries()) {
        console.log(`${pair[0]}: ${pair[1]}`);
      }
    const params = new URLSearchParams(formData);
    console.log(`params: ${params}`)

    // Detect environment and set appropriate base URL
    const isLocalFile = window.location.protocol === 'file:';
    const isLocalhost = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1';
     

    let baseUrl;
    if (isLocalFile || (isLocalhost && window.location.port !== '8080')) {
        // Running locally (file:// or different port)
        baseUrl = 'http://localhost:8080';
    } else {
        // Running on Render or localhost:8080
        baseUrl = window.location.origin;
    }

    let url = `${baseUrl}/api/data?${params.toString()}`
    // let url = `http://localhost:8080/api/data?${params.toString()}`

    console.log(`url sent to backend: ${url}`)

    const response = await fetch(url, {
        method: 'GET',
    })

    const result = await response.json();
    const container = document.getElementById('papers-container');

    container.innerHTML = '';
    result.forEach((paper) => {
    const paperElement = document.createElement('div');
    paperElement.className = 'paper-card';
    paperElement.innerHTML = `
        <h3>${paper.title} 
            <a href=${paper.url} target="_blank">${paper.url}</a>
        </h3>
        <p><strong>Abstract:</strong> ${paper.abstract}</p>
    `;
    container.appendChild(paperElement);
    });
}

init();