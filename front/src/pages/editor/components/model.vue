<template>
	<div>
		<div class="row">
			<h4
				class="col-9"
				data-toggle="collapse"
				v-bind:data-target="'#model_' + model.id"
				v-bind:aria-expanded="!model.id"
				v-bind:aria-controls="'model_' + model.id"
				ref="model"
			>
				{{ model.name || model.title }}
				<i class="fa fa-angle-down"></i>
			</h4>

			<div class="col-3">
				<div class="float-right">
					<button v-if="!model.id" @click="$emit('add', model)">Save</button>
					<button v-if="model.id" class="btn-danger" @click="$emit('remove')">Delete</button>
				</div>
			</div>
		</div>

		<div class="collapse" v-bind:id="'model_' + model.id" v-bind:class="{ show: !model.id }">
			<div class="row mx-0">
				<div class="names pt-4">
					<div>
						<label>English Name</label>
						<input v-model="model.title" />
					</div>
					<div>
						<label>Damage</label>
						<input v-model="model.damage" />
					</div>
					<div>
						<label>Fury/Focus</label>
						<input v-model="model.resource" />
					</div>
					<div>
						<label>Threshold</label>
						<input v-model="model.threshold" />
					</div>
					<div>
						<label>Magic Ability</label>
						<input v-model="model.magic_ability" />
					</div>
				</div>

				<div class="statline">
					<div>
						<span>spd</span>
						<span>str</span>
						<span>mat</span>
						<span>rat</span>
						<span>def</span>
						<span>arm</span>
						<span>cmd</span>
						<span>base</span>
					</div>
					<div>
						<input v-model="model.spd" placeholder="spd" />
						<input v-model="model.str" placeholder="str" />
						<input v-model="model.mat" placeholder="mat" />
						<input v-model="model.rat" placeholder="rat" />
						<input v-model="model.def" placeholder="def" />
						<input v-model="model.arm" placeholder="arm" />
						<input v-model="model.cmd" placeholder="cmd" />
						<input v-model="model.base_size" placeholder="base" />
					</div>
					<label v-for="a in advantages" :key="a.label" v-bind:value="a.label">
						<input type="checkbox" v-model="model.advantages" :value="a.label" />{{ a.name }}
					</label>
				</div>
			</div>

			<div v-if="model.id">
				<div class="row pt-3">
					<div class="col-1"></div>
					<div class="col-11">
						<h4>{{ model.name || model.title }}'s weapons</h4>
						<Weapons :model_id="model.id" />
					</div>
				</div>
			</div>
		</div>
		<hr v-if="model.id" />
	</div>
</template>

<script>
import Weapons from './weapons.vue'
import { ModelAdvantages } from '../../../const.js'
import { EventBus } from '../main.js'

export default {
	name: 'Model',
	props: ['model'],
	components: { Weapons },
	mounted: function() {
		EventBus.$on('mega_save', this.save)
	},
	beforeDestroy() {
		EventBus.$off('mega_save', this.save)
	},
	data() {
		return {
			advantages: ModelAdvantages,
			update: false,
		}
	},
	methods: {
		new: function() {
			if (this.model.id == null) {
				this.model.id = 0
			}
			this.save()
		},
		save: function() {
			if (this.model.id == null) {
				return
			}
			this.$http
				.put(
					process.env.VUE_APP_API_ENDPOINT + '/model/' + this.model.id + '?lang=' + this.$language,
					this.model
				)
				.then(function(res) {
					console.debug(res)
				})
				.catch(function(err) {
					EventBus.$emit('err_save', 'model', this.model.id, err.data)
				})
		},
		open: function() {
			this.$refs.model.click()
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.names {
	@extend .col-4;
	div {
		@extend .row;
		label {
			@extend .col-4;
			@extend .col-form-label;
			@extend .px-0;
		}
		input {
			@extend .col-8;
		}
	}
}
.statline {
	@extend .col-8;
	div {
		@extend .row;
		span {
			@extend .col;
		}
		input {
			@extend .col;
		}
	}
	label {
		@extend .form-check;
		@extend .form-check-inline;
		@extend .form-check-label;
		input {
			@extend .form-check-input;
		}
	}
}
</style>
