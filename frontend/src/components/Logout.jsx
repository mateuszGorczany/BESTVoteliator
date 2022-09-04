import React from 'react';
import { useGoogleLogout } from 'react-google-login';



function Logout() {
	function onLogoutSuccess() {
		alert('Wylogowano pomyślnie')
	}
	function onFailure() {
	}

	const { signOut } = useGoogleLogout(
		{
			clientId: clientID,
			onLogoutSuccess,
			onFailure
		}
	);

	return (
        <button onClick={signOut} className="flex flex-wrap justify-center w-full border border-gray-300 hover:border-gray-500 px-2 py-1.5 rounded-md">
		{/* <button onClick={signIn} className="button"> */}
            <img className="w-5 mr-2" src="https://lh3.googleusercontent.com/COxitqgJr1sJnIDe8-jiKhxDx1FrYbtRHKJ9z_hELisAlapwE9LUPh6fcXIfb5vwpbMl4xl9H9TRFPc5NOO8Sb3VSgIBrfRYvW6cUA" />
			<span className="buttonText">Wyloguj</span>
		</button>
	)
}

export default Logout