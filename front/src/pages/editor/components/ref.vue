<template>
	<div class="ref">
		<div class="header">
			<h2 class="col-8">{{ ref.title }}</h2>
			<div class="col-4">
				<div class="float-right">
					<select v-model="ref.status">
						<option value="wip">WIP</option>
						<option value="tbv">To be validated</option>
						<!-- <option value="done">Done</option> -->
						<option value="tbd">To be deleted</option>
					</select>
					<button v-on:click="save()">Save</button>
					<!-- <button class="btn-danger" v-on:click="remove(ref.id)">Delete</button> -->
				</div>
			</div>
		</div>

		<div class="error" v-if="alert" :class="{ 'alert-success': alert_success, 'alert-danger': !alert_success }">
			<pre class="mb-0">{{ alert }}</pre>
		</div>

		<div class="nav" role="tablist">
			<a class="active" data-toggle="tab" href="#nav-ref">Ref</a>
			<a v-if="ref.id > 0" data-toggle="tab" href="#nav-models">Models</a>
			<a v-if="ref.id > 0" data-toggle="tab" href="#nav-abilities" v-on:click="abilitiesKey++">Abilities</a>
			<a v-if="ref.id > 0" data-toggle="tab" href="#nav-spells" @click="refreshAbilities()">Spells & Animus</a>
			<a
				v-if="ref.id > 0 && [1, 2, 10].includes(ref.category_id)"
				data-toggle="tab"
				href="#nav-feat"
				@click="refreshAbilities()"
				>Feat</a
			>
		</div>

		<div class="content">
			<div class="tab-pane fade show active" id="nav-ref" role="tabpanel" aria-labelledby="nav-ref-tab">
				<Card :reference="ref" :key="ref.id"/>
			</div>
			<div class="tab-pane fade" id="nav-models" role="tabpanel" aria-labelledby="nav-models-tab">
				<Models v-if="ref.id > 0" :ref_id="ref.id" />
			</div>
			<div class="tab-pane fade" id="nav-abilities" role="tabpanel" aria-labelledby="nav-abilities-tab">
				<Abilities v-if="ref.id > 0" :ref_id="ref.id" :key="abilitiesKey"></Abilities>
			</div>
			<div class="tab-pane fade" id="nav-spells" role="tabpanel" aria-labelledby="nav-spells-tab">
				<Spells v-if="ref.id > 0" :ref_id="ref.id" :abilities="abilities"></Spells>
			</div>
			<div class="tab-pane fade" id="nav-feat" role="tabpanel" aria-labelledby="nav-feat-tab">
				<Feat
					v-if="ref.id > 0 && [1, 2, 10].includes(ref.category_id)"
					:abilities="abilities"
					:ref_id="ref.id"
				></Feat>
			</div>
		</div>
	</div>
</template>

<script>
import Abilities from './abilities.vue'
import Models from './models.vue'
import Spells from './spells.vue'
import Feat from './feat.vue'
import Card from './card.vue'
import { EventBus } from '../main.js'
export default {
	name: 'Ref',
	props: ['ref_id'],
	components: { Card, Models, Abilities, Spells, Feat },
	watch: {
		ref_id: function(newVal) {
			this.get(newVal)
		},
	},
	created: function() {
		EventBus.$on('err_save', (component, id, err) => {
			this.alert = this.alert + 'Error in ' + component + ' ' + id + ': ' + err + '\n'
			this.alert_success = false
		})

		this.get(this.ref_id)
	},
	beforeDestroy() {
		EventBus.$off('err_save')
	},
	data() {
		return {
			ref: {},
			alert: '',
			alert_success: false,
			abilitiesKey: 0,
			abilities: [],
		}
	},
	methods: {
		reset: function() {
			this.alert = ''
			this.alert_success = false
		},
		get: function(refID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${refID}?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.ref = res.data
				})
				.catch(function(err) {
					console.error(err)
				})
		},
		save: function() {
			this.reset()
			EventBus.$emit('mega_save')
			this.sleep(1000).then(() => {
				if (this.alert === '') {
					this.alert = 'save success'
					this.alert_success = true
				}
			})
			this.sleep(5000).then(() => {
				if (this.alert_success) {
					this.reset()
				}
			})
		},
		sleep: function(ms) {
			return new Promise(resolve => setTimeout(resolve, ms))
		},
		remove: function(refID) {
			this.alert = ''
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + '/ref/' + refID)
				.then(function(res) {
					console.debug(res)
					// if (res.status === 204) {
					// 	this.$emit('remove_card', card.id)
					// }
				})
				.catch(function(err) {
					this.alert = 'error: ' + err.data
					this.alert_success = false
				})
		},
		refreshAbilities: function() {
			this.$http.get(process.env.VUE_APP_API_ENDPOINT + `/abilities?lang=${this.$language}`).then(function(res) {
				console.debug(res)
				this.abilities = res.data
			})
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.ref {
	.header {
		@extend .row;
		@extend .form-inline;
	}

	.error {
		@extend .row;
		@extend .alert;
		@extend .mx-0;
		@extend .py-1;
	}

	.nav {
		@extend .nav;
		@extend .nav-tabs;

		a {
			@extend .nav-item;
			@extend .nav-link;
		}
	}

	.content {
		@extend .tab-content;
		@extend .container;
		@extend .px-0;
		@extend .mt-4;
	}
}
</style>
