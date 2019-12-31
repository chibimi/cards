<template>
	<div>
		<Model
			v-for="(val, idx) in models"
			v-bind:model="val"
			:key="val.id"
			v-on:remove="removeModel(val, idx)"
			:ref="'model_' + val.id"
		></Model>
		<div class="card">
			<h5 class="header" data-toggle="collapse" data-target="#new_model" aria-expanded="false" ref="newModel">
				New Model
			</h5>
			<div class="collapse card-body p-1" id="new_model">
				<Model v-bind:model="model" v-on:add="addModel"></Model>
			</div>
		</div>
	</div>
</template>

<script>
import Model from './model.vue'
export default {
	name: 'Models',
	props: ['ref_id'],
	components: { Model },
	watch: {
		ref_id: function(newVal) {
			this.get(newVal)
		},
	},
	created: function() {
		this.get(this.ref_id)
	},
	data() {
		return {
			models: [],
			model: {
				ref_id: this.ref_id,
				advantages: [],
				weapons: [],
			},
		}
	},
	methods: {
		get: function(id) {
			this.models = []
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${id}/model?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.models = res.data
				})
		},
		removeModel: function(model, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + `/model/${model.id}`)
				.then(function(res) {
					console.log(res)
					if (res.status === 204) {
						this.models.splice(index, 1)
					}
				})
		},
		addModel: function(model) {
			if (model.id == null) {
				model.id = 0
				model.ref_id = this.ref_id
			}
			this.$http
				.put(
					process.env.VUE_APP_API_ENDPOINT + `/model/${model.id}?lang=${this.$language}`,
					model
				)
				.then(async function(res) {
					console.debug(res)
					if (res.status === 201) {
						model.id = res.data
						this.models.push(model)
						this.model = {
							ref_id: this.ref_id,
							advantages: [],
							weapons: [],
						}
						this.$refs.newModel.click()
						var modelRef = 'model_' + model.id
						while (this.$refs[modelRef] === undefined) {
							await this.sleep(100)
						}
						this.$refs[modelRef][0].open()
					}
				})
		},
		sleep: function(ms) {
			return new Promise(resolve => setTimeout(resolve, ms))
		},
	},
}
</script>

<style scoped></style>
