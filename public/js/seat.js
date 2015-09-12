function escapeHtml(str) {
    str = str.replace(/&/g, '&amp;');
    str = str.replace(/</g, '&lt;');
    str = str.replace(/>/g, '&gt;');
    str = str.replace(/"/g, '&quot;');
    str = str.replace(/'/g, '&#39;');
    return str;
}

function addInputTarrents() {

    var name = document.createElement('input')
    name.type = "text";

    var timing = document.createElement('input')
    timing.type = "text";

    var range = document.createElement('input')
    range.type = "text";

    var dist = document.createElement('input')
    dist.type = "text";

    var cost = document.createElement('input')
    cost.type = "text";

    var effect = document.createElement('input')
    effect.type = "text";
    effect.width = 20;


    var nameTd = document.createElement('td')
    var timingTd = document.createElement('td')
    var rangeTd = document.createElement('td')
    var distTd = document.createElement('td')
    var costTd = document.createElement('td')
    var effectTd = document.createElement('td')

    nameTd.appendChild(name)
    timingTd.appendChild(timing)
    rangeTd.appendChild(range)
    distTd.appendChild(dist)
    costTd.appendChild(cost)
    effectTd.appendChild(effect)

    var elem = document.createElement('tr');

    elem.appendChild(nameTd)
    elem.appendChild(timingTd)
    elem.appendChild(rangeTd)
    elem.appendChild(distTd)
    elem.appendChild(costTd)
    elem.appendChild(effectTd)

    var parent = document.getElementById('tarrents-table');
    parent.appendChild(elem);
}
