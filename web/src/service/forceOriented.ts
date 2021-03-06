import { Link as Edge } from '../model/Link';
import { Node } from '../model/Node';

export class Result{
    constructor(public nodes: Node[], public links: Edge[]){}
}

/**
 * 力导向的配置
 * @export
 * @interface ForceOrientedConfig
 */
export interface ForceOrientedConfig {
    canvasWidth?: number, 
    canvasHeight?: number,
    iterateTimes?: number,
    generateCoordinates?: boolean
}

/**
 * 计算的默认配置
 */
const defaultOption: ForceOrientedConfig = {
    canvasWidth: 1000, 
    canvasHeight: 1000,
    iterateTimes: 200,
    generateCoordinates: true
}

export function forceOriented(mNodeList: Node[], mEdgeList: Edge[], options?: ForceOrientedConfig){
    // 1. 随机分布初始节点位置；
    // 2. 计算每次迭代局部区域内两两节点间的斥力所产生的单位位移（一般为正值）；
    // 3. 计算每次迭代每条边的引力对两端节点所产生的单位位移（一般为负值）；
    // 4. 步骤 2、3 中的斥力和引力系数直接影响到最终态的理想效果，它与节点间的距离、节点在系统所在区域的平均单位区域均有关，需要开发人员在实践中不断调整；
    // 5. 累加经过步骤 2、3 计算得到的所有节点的单位位移；
    // 6. 迭代 n 次，直至达到理想效果。

    options = {...defaultOption, ...options}

    // 常数，如电话常数或者万有引力常数，这里用canvas的尺寸计算
    let k;
    // 记录每一个点x方向的力和
    let mDxMap = {};
    // 记录每一个点y方向的力和
    let mDyMap = {};

    // node的map，增加查询效率
    let mNodeMap = {};

    // 计算k
    if (mNodeList && mEdgeList) {
        k = Math.sqrt(options.canvasWidth * options.canvasHeight / mNodeList.length);
    }

    // 构建mNodeMap
    for (let i = 0; i < mNodeList.length; i++) {
        let node = mNodeList[i];
        if (node) {
            mNodeMap[node.index] = node;
        }
    }

    if(options.generateCoordinates){
        //随机生成坐标. Generate coordinates randomly.
        let initialX, initialY, initialSize = 40.0;
        for (let i in mNodeList) {
            initialX = options.canvasWidth * .5;
            initialY = options.canvasHeight * .5;
            mNodeList[i].x = initialX + initialSize * (Math.random() - .5);
            mNodeList[i].y = initialY + initialSize * (Math.random() - .5);
        }
    }
    
    /**
     * 计算两个Node的斥力产生的单位位移。
     * Calculate the displacement generated by the repulsive force between two nodes.*
     */
    function calculateRepulsive() {
        // 弹出因子？ 不知道干什么的，没有原始公式的文章，应该就是一个修正参数，用于计算距离
        let ejectFactor = 6;
        let distX, distY, dist;
        for (let i = 0; i < mNodeList.length; i++) {
            mDxMap[mNodeList[i].index] = 0.0;
            mDyMap[mNodeList[i].index] = 0.0;
            for (let j = 0; j < mNodeList.length; j++) {
                // 计算i 和 j两点的x差，y差和距离
                if (i !== j) {
                    distX = mNodeList[i].x - mNodeList[j].x;
                    distY = mNodeList[i].y - mNodeList[j].y;
                    dist = Math.sqrt(distX * distX + distY * distY);
                }

                // 如果距离小，弹力因子变小？
                if (dist < 30) {
                    ejectFactor = 5;
                }

                // 距离过远的点忽略不计
                if (dist > 0 && dist < 250) {
                    let id = mNodeList[i].index;
                    // 根电荷公式计算i点的总斥力 k / ( dist^2) / dist * distX。显然和公司不一样？？？
                    mDxMap[id] = mDxMap[id] + distX / dist * k * k / dist * ejectFactor;
                    mDyMap[id] = mDyMap[id] + distY / dist * k * k / dist * ejectFactor;
                }
            }
        }
    }

    /**
     * 计算Edge的引力对两端Node产生的引力。
     * Calculate the traction force generated by the edge acted on the two nodes of its two ends.
     */
    function calculateTraction() {
        let condenseFactor = 3;
        let startNode, endNode;
        for (let e = 0; e < mEdgeList.length; e++) {
            let eStartID = mEdgeList[e].source;
            let eEndID = mEdgeList[e].target;
            startNode = mNodeMap[eStartID];
            endNode = mNodeMap[eEndID];
            if (!startNode) {
                console.log("Cannot find start node id: " + eStartID + ", please check it out.");
                return;
            }
            if (!endNode) {
                console.log("Cannot find end node id: " + eEndID + ", please check it out.");
                return;
            }
            let distX, distY, dist;
            distX = startNode.x - endNode.x;
            distY = startNode.y - endNode.y;
            dist = Math.sqrt(distX * distX + distY * distY);
            // 和我认为的公式也有区别
            mDxMap[eStartID] = mDxMap[eStartID] - distX * dist / k * condenseFactor;
            mDyMap[eStartID] = mDyMap[eStartID] - distY * dist / k * condenseFactor;
            mDxMap[eEndID] = mDxMap[eEndID] + distX * dist / k * condenseFactor;
            mDyMap[eEndID] = mDyMap[eEndID] + distY * dist / k * condenseFactor;
        }
    }

    /**
     * 更新坐标。
     * update the coordinates.
     */
    function updateCoordinates() {
        let maxt = 4, maxty = 3; //Additional coefficients.
        for (let v = 0; v < mNodeList.length; v++) {
            let node = mNodeList[v];
            let dx = Math.floor(mDxMap[node.index]);
            let dy = Math.floor(mDyMap[node.index]);

            if (dx < -maxt) dx = -maxt;
            if (dx > maxt) dx = maxt;
            if (dy < -maxty) dy = -maxty;
            if (dy > maxty) dy = maxty;
            node.x = node.x + dx >= options.canvasWidth || node.x + dx <= 0 ? node.x - dx : node.x + dx;
            node.y = node.y + dy >= options.canvasHeight || node.y + dy <= 0 ? node.y - dy : node.y + dy;
        }
    }

    //迭代200次. Iterate 200 times.
    for (let i = 0; i < options.iterateTimes; i++) {
        calculateRepulsive();
        calculateTraction();
        updateCoordinates();
    }
        
    return new Result(mNodeList, mEdgeList)
}

