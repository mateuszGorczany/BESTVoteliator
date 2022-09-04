const clientID = import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID
const clientSecret = import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_SECRET

function getData(url, data) {
    // const response = await fetch(url, {
    //     method: "GET",
    //     mode: "cors",
    //     cache: "no-cache",
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    // })
    // return response.json()
    return {
        description: "Głosowanie na Fulla",
        options: [
            {
                description: "Adam Abacki", id: 0,
            },
            {
                description: "Paweł Papacki",
                id: 1,
            }
        ],
    }
}

async function postData(url, data) {
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    return response.json()
}

function refreshTokenSetup(res) {
    let refreshTiming = (res.tokenObj.exipres_in || 3600 - 5 * 60) *1000;
    async function refreshToken() {
        const newAuthRes = await res.reloadAuthResponse();
        refreshTiming = (newAuthRes.tokenObj.exipres_in || 3600 - 5 * 60) *1000;
        setTimeout(refreshToken, refreshTiming)
    }
    setTimeout(refreshToken, refreshTiming);
}

export { getData, postData, clientID }