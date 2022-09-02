import { useState } from "react";
import VotingField from "../VotingField"
import { postData } from "../../utils"

const VotingMode = {
    SingleChocie: {
        value: "single",
        description: "Jednokrotny"
    },
    MultipleChocie: {
        value: "multpile",
        description: "Wielokrotny"
    }
}

function sendVotingSubject(votingSubject) {
    console.log(votingSubject)
    const response = postData(import.meta.env.VITE_SERVER_HOST + "/subject", votingSubject)
}

function CreatePage(props) {
    const [choicesValues, setChoices] = useState([])
    const [optionValue, setOptionValue] = useState("")
    const [descriptionValue, setDescriptionValue] = useState("")
    const [votingModeValue, setVotingModeValue] = useState("")
    let choices = choicesValues.map((val, i) => <VotingField key={i}>{val}</VotingField>)

    function onAddBtnClick(event) {
        if (optionValue != "") {
            setChoices(oldArray => [...oldArray, optionValue])
        }
        setOptionValue(() => "")
    };

    function handleSubmit(event) {
        sendVotingSubject({
            descripton: descriptionValue,
            options: choicesValues,
            votingMode: votingModeValue
        })
        setChoices(oldArray => [])
        setDescriptionValue("")
    }
    return (
        <div>
            {/* <form onSubmit={handleSubmit}> */}
            <form >
                <h2>
                    <label>
                        Opis
                    </label>
                </h2>
                <input
                    type="text"
                    name="description"
                    value={descriptionValue}
                    onChange={event => setDescriptionValue(event.target.value)}
                    autoComplete="off"
                    id="voting-subject-description"
                >
                </input>
                <h2>
                    <label>
                        Dodaj pole wyboru
                    </label>
                </h2>
                <input
                    type="text"
                    value={optionValue}
                    onChange={event => setOptionValue(event.target.value)}
                    id="add-option-field"
                    name="text"
                    autoComplete="off"
                />
                <button
                    className="option-button"
                    type="button"
                    onClick={onAddBtnClick}
                >
                    Dodaj
                </button>
                <div className="options-container">
                    <ul className="choices"></ul>
                    {choices}
                </div>
                <select onChange={(event) => setVotingModeValue(event.target.value)}>
                    <option value="single-choice">Jeden wybór</option>
                    <option value="multiple-choice">Wiele wyborów</option>
                </select>
                <button
                    className="submit-button"
                    // type="submit"
                    type="button"
                    onClick={handleSubmit}
                >
                    Prześlij
                </button>
            </form>
        </div >
        // <div class="relative">
        //     <label class="sr-only" for="email"> Opcja </label>
        //     <div>
        //         <input
        //             class="py-4 pl-3 pr-16 text-sm border-2 border-gray-200 rounded-lg"
        //             id="option"
        //             type="option"
        //             placeholder="Option"
        //         />

        //         <button onClick={onAddBtnClick} class="absolute p-2 text-white -translate-y-1/2 bg-blue-600 rounded-full top-1/2 right-4" type="button">
        //             <svg class="w-4 h-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        //                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
        //             </svg>
        //         </button>
        //     </div>
        // </div>
        // <form class="bg-white rounded-md py-10 px-12 shadow-lg">
        // 		<h1 class="text-xl mt-2 text-center font-semibold text-gray-600">Whrite Todo List</h1>
        // 		<div class="mt-6 flex space-x-4 m-10 justify-center">
        // 			<input type="number" placeholder="0" min="0" class="cursor-pointer bg-gray-100 w-10 text-center rounded-md pl-2 outline-none py-2 border-2" />
        // 			<input placeholder="write her..." class="bg-gray-100 rounded-md py-2 px-4 border-2 outline-none"/>
        // 			<button class="bg-yellow-400 px-4 rounded-md font-semibold">Send</button>
        // 		</div>
        // </form>
    )
}

export default CreatePage
