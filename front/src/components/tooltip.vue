<template>
    <div class="tooltip">
        ?
        <span class="tooltiptext">{{txt}}</span>
    </div>
</template>

<script>
export default {
    name: "Tooltip",
    props: ["txt"]
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.tooltip {
  opacity: 1;
  position: relative;
  display: inline-block;
  border-bottom: 1px dotted black;
}

.tooltip .tooltiptext {
  visibility: hidden;
  width: 300px;
  background-color: #555;
  color: #fff;
  text-align: center;
  border-radius: 6px;
  padding: 5px 5px;
  position: absolute;
  z-index: 1;
  bottom: 125%;
  margin-left: -20px;
  transition: opacity 0.3s;
}

.tooltip:hover > .tooltiptext {
  visibility: visible;
  opacity: 1;
}
</style>
