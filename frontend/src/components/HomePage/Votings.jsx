import { getData } from "../../utils"
import VotingSubject from "../VotingSubject"

function getVotings() {
}

function Votings() {
    let votings = getData(import.meta.env.SERVER_HOST + "/subjects")
    // .map((value, i) => <VotingSubject key={i}>{value}</VotingSubject>)
    return (
        <div>
            {votings}
        </div>
    )
}

export default Votings