import { getGoogleUrl } from "../login";

function GoogleButton() {
    return (<div className="mb-3">
        <button href={getGoogleUrl} className="flex flex-wrap justify-center w-full border border-gray-300 hover:border-gray-500 px-2 py-1.5 rounded-md">
            <img className="w-5 mr-2" src="https://lh3.googleusercontent.com/COxitqgJr1sJnIDe8-jiKhxDx1FrYbtRHKJ9z_hELisAlapwE9LUPh6fcXIfb5vwpbMl4xl9H9TRFPc5NOO8Sb3VSgIBrfRYvW6cUA" />
            Zaloguj się
        </button>
    </div>
    )
}

export default GoogleButton