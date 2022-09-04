import React from "react";
import { useGoogleLogin, GoogleLogin, GoogleLogout } from "react-google-login";
import { clientID } from "../utils";

function Login2() {
  function onLoginSuccess(res) {
    alert("Wylogowano pomyślnie");
  }
  function onFailure(res) {}

  const { signIn } = useGoogleLogin({
    onLoginSuccess,
    onFailure,
    clientID,
    isSignedIn: true,
    accessType: "offline",
  });

  return (
    <button
      onClick={signIn}
      className="flex flex-wrap justify-center w-full border border-gray-300 hover:border-gray-500 px-2 py-1.5 rounded-md"
    >
      {/* <button onClick={signIn} className="button"> */}
      <img
        className="w-5 mr-2"
        src="https://lh3.googleusercontent.com/COxitqgJr1sJnIDe8-jiKhxDx1FrYbtRHKJ9z_hELisAlapwE9LUPh6fcXIfb5vwpbMl4xl9H9TRFPc5NOO8Sb3VSgIBrfRYvW6cUA"
      />
      <span className="buttonText">Zaloguj</span>
    </button>
  );
}

function Login() {
  return (
    <div>
      <GoogleLogin
        clientId={clientID}
        buttonText="Zaloguj"
        onSuccess={(response) =>
          console.log("Login success, current user: ", resp.profileObj)
        }
        onFailure={(response) => {
          console.error("Could not login");
          console.log(response);
        }}
        cookiePolicy={"single_host_origin"}
        style={{ marginTop: "100px" }}
        isSignedIn={true}
      />
    </div>
  );
}

function Logout() {
  return (
    <div>
      <GoogleLogout
        clientId={clientID}
        buttonText="Wyloguj"
        onLogoutSuccess={(response) => console.log(response)}
      />
    </div>
  );
}

function GoogleLoginButton() {
  return (
    <div>
      <div
        id="g_id_onload"
        data-client_id={clientID}
        data-context="signin"
        data-ux_mode="popup"
        data-login_uri="http://localhost:3000"
        data-auto_select="true"
        data-itp_support="true"
      ></div>

      <div
        className="g_id_signin"
        data-type="standard"
        data-shape="pill"
        data-theme="outline"
        data-text="signin"
        data-size="large"
        data-logo_alignment="left"
      ></div>
    </div>
  );
}

function GoogleLoginButtonDetailed() {
  return (
    <div>
      <div
        id="g_id_onload"
        data-client_id="1026540081820-jcq9odlpbi214a97fu7ar4furb3njfmf.apps.googleusercontent.com"
        // data-login_uri="http://localhost:3000"
        data-auto_prompt="false"
      ></div>
      <div
        className="g_id_signin"
        data-type="standard"
        data-size="large"
        data-theme="outline"
        data-text="sign_in_with"
        data-shape="rectangular"
        data-logo_alignment="left"
      ></div>
    </div>
  );
}

export { Login, Logout, GoogleLoginButton, GoogleLoginButtonDetailed };
