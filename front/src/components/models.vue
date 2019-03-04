<template>
	<div class="w-100">
		<Model
			v-for="(value,index) in models2"
			v-bind:selectedModel="value"
			:key="value.id"
			v-on:remove="remove(index)"
			v-on:remove_weapon="removeWeapon(index, $event)"
			v-on:add_weapon="addWeapon(index, $event)"
		></Model>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_model"
				aria-expanded="false"
				aria-controls="new_model"
			>New Model</h5>
			<div class="collapse card-body p-1" id="new_model">
				<Model :selectedModel="model" v-on:add="addModel"></Model>
			</div>
		</div>
	</div>
</template>

<script>
import Model from "./model.vue";
export default {
	name: "Models",
	props: ["id", "models"],
	components: {
		Model
	},
	watch: {
		id: function(newVal) {
			this.model.card_id = newVal;
		},
		models: function(newVal) {
			this.models2 = newVal;
		},
		//deep: true
	},
	data() {
		return {
			models2: [],
			model: {
				card_id: this.id,
				advantages: [],
				weapons: []
			}
		};
	},
	methods: {
		removeModel: function(index) {
			this.$emit("remove", index);
		},
		addModel: function(model) {
			this.$emit("add", model);
			this.model = {
				card_id: this.id,
				advantages: [],
				weapons: []
			};
		},
		addWeapon: function(index, weapon) {
			this.$emit("add_weapon", index, weapon);
		},
		removeWeapon: function(index, weaponIndex) {
			this.$emit("remove_weapon", index, weaponIndex);
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
