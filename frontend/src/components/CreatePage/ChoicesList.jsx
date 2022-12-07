import { useState } from "react"
import VotingField from "../VotingField"
import ChoicesForm from "./CreatePage"

const choice = {
    "description": 1.
}

function ChoicesList() {
    return (
        <div>
            <ChoicesForm setChoices={setChoices} onSubmit={onSubmit}></ChoicesForm>
            {choices}
        </div>
    )
}

export default ChoicesList