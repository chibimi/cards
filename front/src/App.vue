<template>
	<div id="app" class="container mt-3">
		<Selector v-on:select_ref="setRef" v-on:change_language="setLanguage" :key="componentKey" />
		<div class="row">
			<Ref v-if="reference.id !== null" :reference="reference" :key="componentKey" />
		</div>
	</div>
</template>

<script>
import Ref from "./components/ref.vue";
import Selector from "./components/selector.vue";
export default {
	name: "app",
	components: {
		Ref,
		Selector
	},
	data() {
		return {
			componentKey: 0,
			ref:null,
			reference: {
				id: null,
			}
		};
	},
	methods: {
		setRef: function(ref){
			if (ref.id === 0&& this.reference.id===0) {
				this.reference.id = -1
			} else {
				this.reference = ref
			}
		},
		setLanguage: function(language) {
			if (language === this.$language) {
				return
			}
			this.$change_language(language)
			this.forceRerender()
		},
		forceRerender() {
			this.componentKey += 1;  
		}
	}
};
</script>

<style>
#app {
	font-family: "Avenir", Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	text-align: center;
	color: #2c3e50;
}

body,
html {
	height: 100%;
}
.h-head {
	height: 10vh;
}
.h-content {
	height: 90vh;
}
</style>
