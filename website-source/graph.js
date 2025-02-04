const width = window.innerWidth;
const height = window.innerHeight;

let links = await d3.json("connections-processed.json")
let nodes= await d3.json("nodes.json")

const svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height)

const simulation = d3.forceSimulation(nodes)
    .force("link", d3.forceLink(links).id(d => d.id).distance(10))
    .force("charge", d3.forceManyBody().strength(-3))
    .force("center", d3.forceCenter(width / 2, height / 2))
    .on("tick", tick);

const link = svg.append("g")
    .attr("stroke", "#eaeaea")
    .attr("stroke-opacity", 0.2)
    .selectAll()
    .data(links)
    .join("line")
    .attr("stroke-width", 0.2);

const node = svg.append("g")
    .selectAll()
    .data(nodes)
    .join("circle")
    .attr("r", 2)
    .attr("fill", d => nodeColor(d.group));

node.append("title")
    .text(d => d.id);

// Add a drag behavior.
node.call(d3.drag()
    .on("start", dragstarted)
    .on("drag", dragged)
    .on("end", dragended));

// Set the position attributes of links and nodes each time the simulation ticks.
function ticked() {
    link
        .attr("x1", d => d.source.x)
        .attr("y1", d => d.source.y)
        .attr("x2", d => d.target.x)
        .attr("y2", d => d.target.y);

    node
        .attr("cx", d => d.x)
        .attr("cy", d => d.y);
}

// Reheat the simulation when drag starts, and fix the subject position.
function dragstarted(event) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    event.subject.fx = event.subject.x;
    event.subject.fy = event.subject.y;
}

// Update the subject (dragged node) position during drag.
function dragged(event) {
    event.subject.fx = event.x;
    event.subject.fy = event.y;
}

// Restore the target alpha so the simulation cools after dragging ends.
// Unfix the subject position now that it’s no longer being dragged.
function dragended(event) {
    if (!event.active) simulation.alphaTarget(0);
    event.subject.fx = null;
    event.subject.fy = null;
}


function tick() {
    link
        .attr("x1", d => d.source.x)
        .attr("y1", d => d.source.y)
        .attr("x2", d => d.target.x)
        .attr("y2", d => d.target.y);

    node
        .attr("cx", d => d.x)
        .attr("cy", d => d.y);
}

container.append(svg.node());

function nodeColor(id) {
    switch(id) {
        case 0: return "#6b1010";
        case 1: return "#1A936F";
        case 2: return "#A734ED";
        case 3: return "#3489ED";
        case 4: return "#C3D24E";
        case 5: return "#ED34CB";
        case 6: return "#43ED34";
        case 7: return "#43ED34";
        case 8: return "#30D392";
        case 9: return "#3E50C7";
        case 10: return "#34EDDE";
        case 11: return "#ED9734";
        case 12: return "#7EA047";
        case 13: return "#72879F";
        case 14: return "#70297E";
        case 15: return "#FC4760";
        case 16: return "#9DA2A8";
        case 17: return "#FFB7FA";
        case 18: return "#E0FFC9";
        case 19: return "#AEB0FF";
        case 20: return "#99ECCB";
        case 21: return "#494C7E";
        case 22: return "#733446";
        case 23: return "#97452C";
    }
}
