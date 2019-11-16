<template>
	<div class="w-100  mt-4" >
		<input v-model="feat.name" type="text" class="form-control" placeholder="Name">
		<textarea v-model="feat.description" type="text" class="form-control mt-2" rows="4" placeholder="Feat description"/>
		<textarea v-model="feat.fluff" type="text" class="form-control mt-2" rows="4" placeholder="Feat fluff (optionnal)"/>
		<div v-if="vo.ref_id" class="col-12 font-italic text-left">
			<p><b>{{vo.name}}:</b> {{vo.description}}</p>
			<p>{{vo.fluff}}</p>
		</div>
		<div class="col-12">
			<div v-if="alert" class="alert alert-error" v-bind:class="{ 'alert-success': alert_success }">{{alert}}</div>
		</div>
	</div>
</template>

<script>
import { EventBus } from '../main.js';

export default {
	name: "Feat",
	props: ["ref_id"],
	watch: {
		ref_id: function(newVal) {
			this.reset();
			this.get(newVal);
		}
	},
	created: function() {
		this.reset();
		this.get(this.ref_id);
	},
	mounted: function(){
		EventBus.$on('mega_save', () => {
			this.save(this.feat)
		})
	},
	data() {
		return {
			vo: {},
			feat: {
				ref_id: this.ref_id
			},
		};
	},
	methods: {
		get: function(refID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/feat?lang=US")
				.then(function(res) {
					console.log(res);
					this.vo = res.data;
				});
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + refID + "/feat?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.feat = res.data;
					if (!this.feat.ref_id) {
						this.feat.ref_id = this.ref_id;
					}
				});
		},
		reset: function() {
			this.alert = "";
			this.alert_succes = false;
		},
		save: function(feat) {
			if (feat.ref_id == null) {
				return
			}
			this.reset();
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + feat.ref_id + "/feat?lang=" + this.$language, feat)
				.then(function(res) {
					console.log(res);
				})
				.catch(function(err) {
				});
		}
	}
};
</script>

<style scoped>
</style>
