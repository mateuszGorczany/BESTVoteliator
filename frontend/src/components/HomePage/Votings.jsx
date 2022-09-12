import { getData } from "../../utils"
import VotingSubject from "../VotingSubject"

function Votings() {
    let votings = getData(import.meta.env.SERVER_HOST + "/subjects")
    return (
        <div>
            {votings}
        </div>
    )
}

export default Votings