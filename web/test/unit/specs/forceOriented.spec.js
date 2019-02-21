import {forceOriented} from '../../../src/service/forceOriented'
import {Link} from '../../../src/model/Link'
import {Node} from '../../../src/model/Node'

describe('forceOriented', () => {
    it('实现力导向布局', () => {
        // 生成一个node列表
        let nodes = []
        let links = []

        for(let i = 0; i < 2000; i++){
            let node = new Node()
            node.index = i
            node.width = 0
            node.height = 0
            nodes.push(node)
        }

        // 分为4组，将各个点连起来，每五个点一次链接，最后4组在各取两个相连
        for(let i = 0; i < 4; i++){
            for(let j = 0; j < 4; j++){
                let link = new Link()
                link.index = i * 5 + j

                link.source = i * 5 + j
                link.target = i * 5 + j + 1

                links.push(link)
            }
        }
        for(let i = 0; i < 3; i++){
            let link = new Link()
            link.index = i * 5 + 5

            link.source = i * 5
            link.target = (i + 1) * 5 + 4 - i

            links.push(link)
        }
        
        forceOriented(nodes, links, {
            iterateTimes: 10000
        })

        let canvas = document.createElement('canvas')
        let context = canvas.getContext('2d')
        canvas.width =1000
        canvas.height = 1000

        context.beginPath()
        let map = {}
        for(let i = 0; i < nodes.length; i++){
            map[nodes[i].index] = nodes[i]
            context.rect(nodes[i].x, nodes[i].y, 10, 10)
        }

        for(let i = 0; i < links.length; i++){
            let link = links[i]
            context.moveTo(map[link.source].x, map[link.source].y)
            context.lineTo(map[link.target].x, map[link.target].y)
        }

        context.fill()
        context.stroke()
        context.closePath()

        let dataURL = canvas.toDataURL()

        console.log(dataURL)

        window.top.document.body.appendChild(canvas)

    })
})
