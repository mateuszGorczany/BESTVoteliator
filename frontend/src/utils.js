const clientID = import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID
const clientSecret = import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_SECRET
const ServerHost = import.meta.env.VITE_SERVER_HOST

async function getData(endpoint, setter) {
    fetch(import.meta.env.VITE_SERVER_HOST + endpoint, {
        method: "GET",
        mode: "cors",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem("GoogleJWT"),
        }
    }).then((response) => response.json())
        .then((response) => setter(response))
        .catch((err) => console.log(err));
}

async function postData(url, data) {
    const response = await fetch(url, {
        method: "POST",
        mode: "cors",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem("GoogleJWT"),
        },
        body: JSON.stringify(data)
    }).then((response) => response.json()).catch((err) => console.log(err))
    console.log(response)
    return response
}

function refreshTokenSetup(res) {
    let refreshTiming = (res.tokenObj.exipres_in || 3600 - 5 * 60) * 1000;
    async function refreshToken() {
        const newAuthRes = await res.reloadAuthResponse();
        refreshTiming = (newAuthRes.tokenObj.exipres_in || 3600 - 5 * 60) * 1000;
        setTimeout(refreshToken, refreshTiming)
    }
    setTimeout(refreshToken, refreshTiming);
}

function isLoggedIn() {
    return localStorage.getItem("GoogleJWT") != null
}

function handleCredentialResponse(response) {
    console.log("Encoded JWT ID token: " + response.credential);
    localStorage.removeItem("GoogleJWT")
    localStorage.setItem("GoogleJWT", response.credential)
    postData(
        ServerHost + "/api/v1/login",
        {
            "GoogleJWT": response.credential
        }
    )
}

function initLogin() {
    google.accounts.id.initialize({
        client_id: import.meta.env.VITE_GOOGLE_OAUTH_CLIENT_ID,
        callback: handleCredentialResponse
    });
    google.accounts.id.renderButton(
        document.getElementById("signInDiv"),
        { theme: "outline", size: "large", shape: "pill", locale: "pl_PL" }  // customization attributes
    );
    google.accounts.id.prompt(); // also display the One Tap dialog
}

export { getData, postData, clientID, initLogin }