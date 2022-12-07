
import { getData } from '../../utils'
import VotingSubject from '../VotingSubject'
import { useState, useEffect } from 'react'

function Home() {
    const [polls, setPolls] = useState([])
    useEffect(() => {
        getData("/api/v1/poll/all", setPolls)
    }, [])

    console.log(polls)
    let pollsElements = polls.map((val, i) => <VotingSubject key={i} title={val.name}>{val.description}</VotingSubject>)
    console.log(pollsElements)
    return (
        <div>
            <h1 className="text-3xl font-bold underline">
                Utworzone głosowania
            </h1>
            <div className="flex">
                <div className='m-auto'>
                    {pollsElements}
                </div>
            </div>
        </div>
    )
}

export default Home