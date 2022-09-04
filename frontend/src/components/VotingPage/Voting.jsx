import { getData } from "../../utils"
import VotingSubject from "../VotingSubject"

function Choice(props) {
	return (
		<div>
		<label><input type="radio" name={props.id} required/>{props.children}</label>
		</div>
	)
}

function VotingForm() {
	let id = 1
	let voting = getData(import.meta.env.SERVER_HOST + "/subjects/" + id)
	let choices = voting.options.map((value, i) => <Choice key={i} name={value.id}>{value.description}</Choice>)
	return (
	<form>
		<label>		{voting.description}</label>
		{choices}
		<button
			className="submit-button"
			// type="submit"
			type="button"
			// onClick={handleSubmit}
		>
			Prześlij
		</button>
	</form>
	)
}

function Voting() {
	// console.log(voting.options)
	console.log(choices)
    return (
        <div>
			<VotingForm/>
        </div>
    )
}

export default Voting