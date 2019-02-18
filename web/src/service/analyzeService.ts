import axios from 'axios'
import { Node } from '../model/Node';
import { Link } from '../model/Link';
import { Group } from '../model/Group';

export function analyze(path: string): Promise<{
    nodes: Node[],
    links: Link[]
    groups: Group[]
}>{
    return axios.get(`analyze?path=${path}`).then(data=> data.data)
}