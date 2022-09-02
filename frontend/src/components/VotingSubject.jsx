import Tag from "./Tag"

function VotingSubject(props) {
    console.log(props)
    let tagsValues = [
        "#travel",
        "#winter",
        "#money"
    ]

    let tags = tagsValues.map((val, i) => <Tag key={i}>{val}</Tag>)
    console.log(tags)
    return (
        <div className="max-w-sm rounded overflow-hidden shadow-lg">
            <div className="px-6 py-4">
                <div className="font-bold text-xl mb-2">The Coldest Sunset</div>
                <p className="text-gray-700 text-base">
                    {props.children}
                </p>
            </div>
            <div className="px-6 pt-4 pb-2">
                {tags}
            </div>
        </div>
    )
}

export default VotingSubject
