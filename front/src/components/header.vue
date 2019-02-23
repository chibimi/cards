<template>
	<div class="dark">{{name}} is coming</div>
</template>

<script>
export default {
	name: "Header",
	props: {
		name: String
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.dark {
	color: #17a161;
}
</style>
