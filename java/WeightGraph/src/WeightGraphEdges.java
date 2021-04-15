public class WeightGraphEdges implements Comparable<WeightGraphEdges> {
    private int v, w, weight;
    public WeightGraphEdges(int v, int w, int weight) {
        this.v = v;
        this.w = w;
        this.weight = weight;
    }

    public int V() {
        return v;
    }

    public int W() {
        return w;
    }

    public int weight() {
        return weight;
    }

    @Override
    public int compareTo(WeightGraphEdges o) {
        return weight - o.weight;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("("+v+"-"+w+":"+weight+")");
        return sb.toString();
    }
}