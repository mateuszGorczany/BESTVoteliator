import Tag from "./Tag"

function VotingSubject(props) {
    let tags = null
    if (props.tags != null) {
        tags = props.tags.map((val, i) => <Tag key={i}>{val}</Tag>)
    }
    return (
        <div className="max-w-sm rounded overflow-hidden shadow-lg ">
            <div className="px-6 py-4">
                <div className="font-bold text-xl mb-2">{props.title}</div>
                <p className="text-gray-700 text-base">
                    {props.children}
                </p>
            </div>
            <div className="px-6 pt-4 pb-2">
                {tags || <div></div>}
            </div>
        </div>
    )
}

export default VotingSubject
