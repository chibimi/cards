<template>
	<div id="app" class="container">
		<div class="row">
			<form class="form-inline">
				<label >Faction</label>
				<select v-model="faction" class="form-control" @change="getModels(faction, category)">
					<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
				</select>
				<label >Category</label>
				<select v-model="category" class="form-control" @change="getModels(faction, category)">
					<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
				</select>
				<label >Model</label>
				<select v-model="model" class="form-control">
					<option value="0">Create new</option>
					<option v-for="m in models" :key="m.id" :value="m.id">[{{m.status}}] {{m.name}}</option>
				</select>
				<button type="submit" class="btn btn-primary">Go</button>
			</form>
		</div>
		<div class="row">
			<Content v-if="model" :id="model" :factions="factions" :categories="categories"/>
		</div>
	</div>
</template>

<script>
import Content from "./components/content.vue"
export default {
	name: "app",
	components: {
		Content
	},
	methods: {
		getModels: function(faction, category) {
			if (!faction || !category){
				return
			}
			this.$http
			.get("http://localhost:9901/models?faction="+faction+"&category="+category)
			.then(function(res) {
				console.log(res);
				this.models = res.data;
			})
			.catch(function(err) {
				console.log(err);
			});
		}
	},
	created: function() {
		this.$http
			.get("http://localhost:9901/factions")
			.then(function(res) {
				console.log(res);
				this.factions = res.data;
			})
			.catch(function(err) {
				console.log(err);
			});
		this.$http
			.get("http://localhost:9901/categories")
			.then(function(res) {
				console.log(res);
				this.categories = res.data;
			})
			.catch(function(err) {
				console.log(err);
			});
	},
	data() {
		return {
			factions: [],
			faction: null,
			categories: [],
			category: null,
			models: [],
			model: null
		};
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
