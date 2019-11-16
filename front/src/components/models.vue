<template>
	<div class="w-100">
		<Model
			v-for="(value,index) in models2"
			v-bind:model="value"
			:key="value.id"
			v-on:remove="removeModel(index)"
			:ref="'model_' + value.id"
		></Model>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_model"
				aria-expanded="false"
				aria-controls="new_model"
				ref="newModel"
			>New Model</h5>
			<div class="collapse card-body p-1" id="new_model">
				<Model :model="newModel" v-on:add="addModel"></Model>
			</div>
		</div>
	</div>
</template>

<script>
import Model from "./model.vue";
export default {
	name: "Models",
	props: ["ref_id"],
	components: {
		Model
	},
	watch: {
		ref_id: function(newVal) {
			this.get(newVal);
		},
	},
	created: function() {
		this.get(this.ref_id);
	},
	data() {
		return {
			models2: [],
			newModel: {
				ref_id: this.ref_id,
				advantages: [],
				weapons: []
			}
		};
	},
	methods: {
		get: function(id) {
			this.models2 = [];
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + id + "/model?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.models2 = res.data;
				});
		},
		removeModel: function(index) {
			this.models2.splice(index, 1);
		},
		addModel: async function(model) {
			this.models2.push(model);
			this.newModel = {
				ref_id: this.id,
				advantages: [],
				weapons: []
			};
			this.$refs.newModel.click();
			
			var modelRef = "model_"+model.id;
			while(this.$refs[modelRef] === undefined){
				await this.sleep(100)
			}
			this.$refs[modelRef][0].open();

		},
		sleep: function(ms) {
			return new Promise(resolve => setTimeout(resolve, ms));
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
