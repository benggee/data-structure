public class BridgeEdge {
    private int v,w;
    
    public BridgeEdge(int v, int w) {
        this.v = v;
        this.w = w;
    }

    @Override
    public String toString() {
        return String.format("%d-%d", v, w);
    }
}