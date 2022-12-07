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
    const response = postData(import.meta.env.VITE_SERVER_HOST + "/api/v1/poll", votingSubject)
}

function CreatePage(props) {
    const [choicesValues, setChoices] = useState([])
    const [optionValue, setOptionValue] = useState("")
    const [descriptionValue, setDescriptionValue] = useState("")
    const [nameValue, setNameValue] = useState("")
    const [votingModeValue, setVotingModeValue] = useState("")
    let choices = choicesValues.map((val, i) => <VotingField key={i}>{val}</VotingField>)

    function onAddBtnClick(event) {
        if (optionValue != "") {
            setChoices(oldArray => [...oldArray, optionValue])
        }
        setOptionValue(() => "")
    };

    function handleSubmit(event) {
        let poll = {
            name: nameValue,
            description: descriptionValue,
            options: choicesValues.map((value, idx, arr) => arr[idx] = {
                description: value,
                option_id: idx
            }),
            votingMode: votingModeValue
        }
        sendVotingSubject(poll)
        setChoices(oldArray => [])
        setDescriptionValue("")
        setNameValue("")
    }
    return (
        <div>
            <form >
                <h2>
                    <label>
                        Nazwa
                    </label>
                </h2>
                <input
                    type="text"
                    name="name"
                    value={nameValue}
                    onChange={event => setNameValue(event.target.value)}
                    autoComplete="off"
                    id="voting-subject-name"
                >
                </input>
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
                    type="button"
                    onClick={handleSubmit}
                >
                    Prześlij
                </button>
            </form>
        </div >
    )
}

export default CreatePage
