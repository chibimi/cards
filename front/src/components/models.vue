<template>
	<div class="w-100">
		<Model v-for="(value,index) in models" v-bind:model="value" :key="value.id" v-on:remove="removeModel(index)"></Model>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_model"
				aria-expanded="false"
				aria-controls="new_model"
			>New Model</h5>
			<div class="collapse card-body p-1" id="new_model">
				<Model :model="model" v-on:add="addModel"></Model>
			</div>
		</div>
	</div>
</template>

<script>
import Model from "./model.vue";
export default {
	name: "Models",
	props: ["id"],
	components: {
		Model
	},
	watch: {
		id: function(newVal, oldVal) {
			this.get(newVal)
		}
	},
	created: function() {
		this.get(this.id)
	},
	data() {
		return {
			models: [],
			model: {
				card_id: this.id,
				advantages: []
			}
		};
	},
	methods: {
		get: function(cardID) {
			this.$http
				.get("http://localhost:9901/cards/" + cardID + "/models")
				.then(function(res) {
					this.models = res.data;
				});
		},
		removeModel: function(index) {
			this.models.splice(index, 1);
		},
		addModel: function(model) {
			this.models.push(model);
			this.model = {
				card_id: this.id, 
				advantages: [] 
			};
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
